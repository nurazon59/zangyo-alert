module SocialInsurance

open FSharp.Data
open System.IO

type Grades = CsvProvider<"data/social-insurance-tokyo-2026.csv">

let private path =
    Path.Combine(__SOURCE_DIRECTORY__, "data", "social-insurance-tokyo-2026.csv")

let resolveGrades (averageAmount: decimal) : decimal =
    Grades.Load(path).Rows
    |> Seq.find (fun row ->
        [ row.Lower
              |> Option.ofNullable
              |> Option.forall (fun lower -> decimal lower < averageAmount)
          row.Upper
              |> Option.ofNullable
              |> Option.forall (fun upper -> averageAmount < decimal upper) ]
        |> List.forall id)
    |> fun row -> decimal row.Amount
