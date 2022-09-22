package status

import (
	"fmt"
	"sync"

	"github.com/BF-Moritz/monitoring.go/api/enum"
	"github.com/BF-Moritz/monitoring.go/api/types"
	"github.com/BF-Moritz/monitoring.go/api/vars"
	"github.com/BF-Moritz/utils.go/math"
)

const maxConcurrentParsers = 10
const numberOfFields = 5

func (s *Service) GetAll(lastMinutes uint32) (status []types.StatusType, err error) {
	unparsedStatus, err := s.StatusDAO.GetAll(lastMinutes)
	if err != nil {
		vars.Logger.LogError("service:status.GetAll()", "failed to get all (%s)", err)
		return nil, err
	}

	status = make([]types.StatusType, 0, len(unparsedStatus))

	var distChan chan types.ParseInType = make(chan types.ParseInType)
	var gathererChan chan types.ParseOutType = make(chan types.ParseOutType)
	var parserWG sync.WaitGroup = sync.WaitGroup{}
	var gathererWG sync.WaitGroup = sync.WaitGroup{}

	var concurrentErrors []error = make([]error, 0)

	var statusMap map[string]types.StatusType = make(map[string]types.StatusType)

	// start parser
	for i := 0; i < math.Min(len(unparsedStatus)*numberOfFields, maxConcurrentParsers); i++ {
		go s.parser(&distChan, &gathererChan, &parserWG)
		parserWG.Add(1)
	}

	// start gatherer
	go s.gatherer(&gathererChan, &gathererWG, &concurrentErrors, &statusMap)
	gathererWG.Add(1)

	// send parse requests
	for _, unparsedData := range unparsedStatus {
		distChan <- types.ParseInType{
			Client: unparsedData.Client,
			Type:   enum.FloatParseType,
			Values: unparsedData.CPUUsage,
			Field:  enum.FieldCPUUsage,
		}
		distChan <- types.ParseInType{
			Client: unparsedData.Client,
			Type:   enum.FloatParseType,
			Values: unparsedData.MemoryCached,
			Field:  enum.FieldMemoryCached,
		}
		distChan <- types.ParseInType{
			Client: unparsedData.Client,
			Type:   enum.FloatParseType,
			Values: unparsedData.MemoryFree,
			Field:  enum.FieldMemoryFree,
		}
		distChan <- types.ParseInType{
			Client: unparsedData.Client,
			Type:   enum.FloatParseType,
			Values: unparsedData.MemoryUsed,
			Field:  enum.FieldMemoryUsed,
		}
		distChan <- types.ParseInType{
			Client: unparsedData.Client,
			Type:   enum.UintParseType,
			Values: unparsedData.MemoryTotal,
			Field:  enum.FieldMemoryTotal,
		}
	}

	close(distChan)
	parserWG.Wait()
	close(gathererChan)
	gathererWG.Wait()

	if len(concurrentErrors) > 0 {
		for i, concurrentError := range concurrentErrors {
			vars.Logger.LogError("service:status.GetAll()", "concurrent error %d/%d: (%s)", i+1, len(concurrentErrors), concurrentError)
		}
		return nil, fmt.Errorf("%d concurrent errors", len(concurrentErrors))
	}

	for _, entry := range statusMap {
		status = append(status, entry)
	}

	return status, nil
}

func (s *Service) parser(inChan *chan types.ParseInType, outChan *chan types.ParseOutType, wg *sync.WaitGroup) {
Global:
	for entry := range *inChan {

		var outEntry = types.ParseOutType{Uints: nil, Floats: nil, Err: nil, Field: entry.Field, Client: entry.Client}

		switch entry.Type {
		case enum.UintParseType:
			var outValues []types.UintDataPoint = make([]types.UintDataPoint, 0, len(entry.Values))
			for _, val := range entry.Values {
				parsed, err := val.ParseToUintDataPoint()
				if err != nil {
					vars.Logger.LogError("service:status.parser()", "failed to parse uint data point (%s)", err)
					outEntry.Err = err
					*outChan <- outEntry
					continue Global
				}

				outValues = append(outValues, parsed)
			}
			outEntry.Uints = &outValues

		case enum.FloatParseType:
			var outValues []types.FloatDataPoint = make([]types.FloatDataPoint, 0, len(entry.Values))
			for _, val := range entry.Values {
				parsed, err := val.ParseToFloatDataPoint()
				if err != nil {
					vars.Logger.LogError("service:status.parser()", "failed to parse float data point (%s)", err)
					outEntry.Err = err
					*outChan <- outEntry
					continue Global
				}

				outValues = append(outValues, parsed)
			}
			outEntry.Floats = &outValues
		}
		*outChan <- outEntry
	}
	wg.Done()
}

func (s *Service) gatherer(inChan *chan types.ParseOutType, wg *sync.WaitGroup, concurrentErrors *[]error, statusMap *map[string]types.StatusType) {
	for entry := range *inChan {
		if entry.Err != nil {
			*concurrentErrors = append(*concurrentErrors, entry.Err)
			vars.Logger.LogError("service:status.gatherer()", "new concurrent error (%s)", entry.Err)
			continue
		}

		if _, ok := (*statusMap)[entry.Client]; !ok {
			(*statusMap)[entry.Client] = types.StatusType{
				Client:       entry.Client,
				CPUUsage:     []types.FloatDataPoint{},
				MemoryTotal:  []types.UintDataPoint{},
				MemoryUsed:   []types.FloatDataPoint{},
				MemoryCached: []types.FloatDataPoint{},
				MemoryFree:   []types.FloatDataPoint{},
			}
		}

		v := (*statusMap)[entry.Client]

		switch entry.Field {
		case enum.FieldCPUUsage:
			if entry.Floats != nil {
				v.CPUUsage = *entry.Floats
			}
		case enum.FieldMemoryCached:
			if entry.Floats != nil {
				v.MemoryCached = *entry.Floats
			}
		case enum.FieldMemoryFree:
			if entry.Floats != nil {
				v.MemoryFree = *entry.Floats
			}
		case enum.FieldMemoryTotal:
			if entry.Uints != nil {
				v.MemoryTotal = *entry.Uints
			}
		case enum.FieldMemoryUsed:
			if entry.Floats != nil {
				v.MemoryUsed = *entry.Floats
			}
		}

		(*statusMap)[entry.Client] = v
	}
	wg.Done()
}

func (s *Service) GetByName(name string) (status []types.StatusType, err error) {
	// TODO
	return []types.StatusType{}, nil
}
