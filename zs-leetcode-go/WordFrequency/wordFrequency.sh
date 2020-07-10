# cat words.txt | awk '{for(i=1;i<=NF;i++) print $i}' | sort -nr | uniq -c | awk '{print $2,$1}'
tr -s ' ' '\n' < words.txt | sort | uniq -c | sort -nr | awk '{print $2, $1}'
