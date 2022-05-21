
const searchOrigin = 'https://mintz5.duckdns.org/bible/search?q='
// const searchOrigin = 'http://localhost:8000/v1/search?q='


function searchforit(){
    // clear scripture area
    document.getElementById('scripture').innerHTML = '';

    // let scriptureAreaHandle = document.getElementById('scripture').querySelectorAll('p');
    // scriptureAreaHandle.forEach(function (item){
    //     item.remove();
    // }) 
    console.log("cleared scipture area..");
    console.log("searching for it");
    console.log("search url: ", searchOrigin + myTextInput.value);
    fetch(searchOrigin + myTextInput.value)
    .then(response => {
        if(!response.ok){
            throw new Error(`HTTP Error: ${response.status}`);
        }
        return response.json();
    }
    )
    .then(function (data){
        for(i=0; i<data.length;i++){
            console.log(data[i]);
            let a = document.createElement('a');
            a.textContent = data[i].book + " " + data[i].chapter + ":" + data[i].verse;
            a.href = "bible" + "/" + data[i].book + '/' + data[i].chapter;
            console.log(a);
            let l = document.createElement('p');
            
            l.textContent = data[i].text;
            l.appendChild(a);
            let p = document.createElement('span');
            p.appendChild(a);
            p.appendChild(l);
           
            document.getElementById('scripture').appendChild(p);
        }
    })
}
