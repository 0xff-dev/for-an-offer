select a.Id from Weather a, Weather b where datediff(a.RecordDate, b.RecordDate) = 1 and a.Temperature > b.Temperature
