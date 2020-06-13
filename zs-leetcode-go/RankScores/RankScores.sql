-- 这个解法是通过了的，注意Rank是mysql的关键字即可。
select A.Score as score, count(*) as `Rank` from Scores A join (select distinct Score from Scores) as B on B.Score >= A.Score group by A.Id, A.Score order by A.Score desc;

-- 下面返回结果如下, 但是leetcode报错
+-------+------+
| Score | Rank |
+-------+------+
|  4.00 |    1 |
|  4.00 |    1 |
|  3.85 |    2 |
|  3.65 |    3 |
|  3.65 |    3 |
|  3.50 |    4 |
+-------+------+

select Scores.Score, rank.r as `Rank` from
    Scores left join (select Score, (@row:=@row+1) as r from (select distinct Score from Scores order by Score desc) as t1, (select @row:=0) as t2) as rank
    on Scores.Score=rank.Score order by Scores.Score desc;