module Config

open System.IO
open FSharp.Data

type ConfigJson =
    JsonProvider<"""
    {
      "BaseSalary": 300000,
      "FixedOvertime": 30,
      "FixedOverSalary": 50000,
      "FixedBenefits": {
        "Housing": 30000,
        "Transport": 10000
      }
    }
    """>

let Config =
    Path.Combine(__SOURCE_DIRECTORY__, "testdata", "sample.json") |> ConfigJson.Load
