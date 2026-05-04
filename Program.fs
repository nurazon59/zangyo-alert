module Program

open System.IO
open FSharp.Data

type Sample = JsonProvider<""" {"name":"x"} """>

do
    Path.Combine(__SOURCE_DIRECTORY__, "testdata", "sample.json")
    |> Sample.Load
    |> _.Name
    |> printfn "%s"

do
    SocialInsurance.printGrades 100000m
