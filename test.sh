for d in ./src/*; do
    for subd in $d/*; do 
        go test ./$subd
        echo $subd
    done
done