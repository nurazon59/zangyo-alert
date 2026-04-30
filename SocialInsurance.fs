module SocialInsurance

open FSharp.Data
open System.IO

let printGrades () =
    Path.Combine(__SOURCE_DIRECTORY__, "data", "social-insurance-tokyo-2026.csv")
    |> CsvFile.Load
    |> _.Rows
    |> Seq.iter (fun row -> printfn "%s" row.["amount"])
