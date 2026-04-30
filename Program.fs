open System.IO
open FSharp.Data

type Sample = JsonProvider<""" {"name":"x"} """>

Path.Combine(__SOURCE_DIRECTORY__, "testdata", "sample.json")
|> Sample.Load
|> _.Name
|> printfn "%s"
