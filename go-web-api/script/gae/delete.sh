VERSIONS=$(gcloud app versions list --project "$1" --service "$2" --sort-by '~version' --format 'value(version.id)' | sort -r)
COUNT=0
for VERSION in $VERSIONS
do
    ((COUNT++))
    if [ $COUNT -gt "$3" ]
    then
        echo "削除 : $VERSION ($1 : $2)"
        gcloud app versions delete "$VERSION" --project "$1" --service "$2" -q
    else
        echo "維持 : $VERSION ($1 : $2)"
    fi
done
