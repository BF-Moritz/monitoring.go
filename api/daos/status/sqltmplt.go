package status

const getAllSQL = `
from(bucket: "monitoring")
  |> range(start: -%dm) 
  |> filter(fn: (r) => r["_measurement"] == "cpu" or r["_measurement"] == "memory")
`
