select 
    CONCAT('https://',requestAddr,request_uri) AS "请求地址", 
    request_uri, 
    totalCount as "请求数",
    account,
    requestId
        from (
            select 
                "content.RequestAddr" as requestAddr,
                case when strpos("content.RequestPath", '?') > 0 then split_part("content.RequestPath", '?', 1) else "content.RequestPath" end as request_uri,
                COUNT("content.RequestPath") as totalCount,
                max_by("content.downstream_X-Request-Id","content.Duration") requestId,
                max_by("content.request_X-Auth-Accountid","content.Duration") account,
                where "content.RequestAddr" like '%tx%'
                group by "content.RequestAddr", request_uri) t
    where totalCount <= 1000 order by totalCount asc limit 100 


(* NOT content.RequestPath:internal*) | select CONCAT('https://',requestAddr,request_uri) AS "请求地址",request_uri, (cast(itemCount as double )/cast(totalCount as double )) as ratio,itemCount as "慢接口请求数",totalCount as "请求数",account,requestId,costTime as "最慢响应时间"  from ( select "content.RequestAddr" as requestAddr, case when strpos("content.RequestPath", '?') > 0 then split_part("content.RequestPath", '?', 1) else "content.RequestPath" end as request_uri ,COUNT("content.RequestPath") as totalCount, count_if(cast("content.Duration" AS bigint)  > 1000000000) as itemCount,max_by("content.downstream_X-Request-Id","content.Duration") requestId,max_by("content.request_X-Auth-Accountid","content.Duration") account, cast(max("content.Duration") as double)/1000000000 as costTime where "content.RequestAddr" like '%tx%' group by "content.RequestAddr",request_uri  order by itemCount desc) t  where totalCount > 500 order by ratio desc limit 100 


(* NOT content.RequestPath:internal* AND content.request_Content-Type:"application/json") | select CONCAT('https://',requestAddr,request_uri) AS "请求地址", request_uri, totalCount as "请求数", account, requestId from (select "content.RequestAddr" as requestAddr,case when strpos("content.RequestPath", '?') > 0 then split_part("content.RequestPath", '?', 1) else "content.RequestPath" end as request_uri,COUNT("content.RequestPath") as totalCount,max_by("content.downstream_X-Request-Id","content.RequestPath") requestId,max_by("content.request_X-Auth-Accountid","content.RequestPath") account where "content.RequestAddr" like '%tx%' group by "content.RequestAddr", request_uri) t where totalCount <= 1000 order by totalCount asc limit 100 