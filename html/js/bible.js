const books = [{'book_name': 'GENESIS', 'ordinal_verse_start': 1508, 'chapters': 50}, {'book_name': 'EXODUS', 'ordinal_verse_start': 2709, 'chapters': 40}, {'book_name': 'LEVITICUS', 'ordinal_verse_start': 3572, 'chapters': 27}, {'book_name': 'NUMBERS', 'ordinal_verse_start': 4881, 'chapters': 36}, {'book_name': 'DEUTERONOMY', 'ordinal_verse_start': 5841, 'chapters': 34}, {'book_name': 'JOSHUA', 'ordinal_verse_start': 6478, 'chapters': 24}, {'book_name': 'JUDGES', 'ordinal_verse_start': 7104, 'chapters': 21}, {'book_name': 'RUTH', 'ordinal_verse_start': 7192, 'chapters': 4}, {'book_name': '1SAMUEL', 'ordinal_verse_start': 8011, 'chapters': 31}, {'book_name': '2SAMUEL', 'ordinal_verse_start': 8694, 'chapters': 24}, {'book_name': '1KINGS', 'ordinal_verse_start': 9482, 'chapters': 22}, {'book_name': '2KINGS', 'ordinal_verse_start': 10224, 'chapters': 25}, {'book_name': '1CHRONICLES', 'ordinal_verse_start': 11166, 'chapters': 29}, {'book_name': '2CHRONICLES', 'ordinal_verse_start': 11995, 'chapters': 36}, {'book_name': 'EZRA', 'ordinal_verse_start': 12254, 'chapters': 10}, {'book_name': 'NEHEMIAH', 'ordinal_verse_start': 12673, 'chapters': 13}, {'book_name': 'ESTHER', 'ordinal_verse_start': 12868, 'chapters': 10}, {'book_name': 'JOB', 'ordinal_verse_start': 13924, 'chapters': 42}, {'book_name': 'PSALMS', 'ordinal_verse_start': 16396, 'chapters': 150}, {'book_name': 'PROVERBS', 'ordinal_verse_start': 17286, 'chapters': 31}, {'book_name': 'ECCLESIASTES', 'ordinal_verse_start': 17525, 'chapters': 12}, {'book_name': 'SONG OF SOLOMON', 'ordinal_verse_start': 17642, 'chapters': 8}, {'book_name': 'ISAIAH', 'ordinal_verse_start': 18924, 'chapters': 66}, {'book_name': 'JEREMIAH', 'ordinal_verse_start': 20278, 'chapters': 52}, {'book_name': 'LAMENTATIONS', 'ordinal_verse_start': 20444, 'chapters': 5}, {'book_name': 'EZEKIEL', 'ordinal_verse_start': 21704, 'chapters': 48}, {'book_name': 'DANIEL', 'ordinal_verse_start': 22083, 'chapters': 12}, {'book_name': 'HOSEA', 'ordinal_verse_start': 22284, 'chapters': 14}, {'book_name': 'JOEL', 'ordinal_verse_start': 22345, 'chapters': 3}, {'book_name': 'AMOS', 'ordinal_verse_start': 22497, 'chapters': 9}, {'book_name': 'OBADIAH', 'ordinal_verse_start': 22512, 'chapters': 1}, {'book_name': 'JONAH', 'ordinal_verse_start': 22570, 'chapters': 4}, {'book_name': 'MICAH', 'ordinal_verse_start': 22666, 'chapters': 7}, {'book_name': 'NAHUM', 'ordinal_verse_start': 22714, 'chapters': 3}, {'book_name': 'HABAKKUK', 'ordinal_verse_start': 22770, 'chapters': 3}, {'book_name': 'ZEPHANIAH', 'ordinal_verse_start': 22822, 'chapters': 3}, {'book_name': 'HAGGAI', 'ordinal_verse_start': 22857, 'chapters': 2}, {'book_name': 'ZECHARIAH', 'ordinal_verse_start': 23070, 'chapters': 14}, {'book_name': 'MALACHI', 'ordinal_verse_start': 23140, 'chapters': 4}, {'book_name': 'MATTHEW', 'ordinal_verse_start': 24197, 'chapters': 28}, {'book_name': 'MARK', 'ordinal_verse_start': 24875, 'chapters': 16}, {'book_name': 'LUKE', 'ordinal_verse_start': 25993, 'chapters': 24}, {'book_name': 'JOHN', 'ordinal_verse_start': 26900, 'chapters': 21}, {'book_name': 'ACTS', 'ordinal_verse_start': 27901, 'chapters': 28}, {'book_name': 'ROMANS', 'ordinal_verse_start': 28338, 'chapters': 16}, {'book_name': '1CORINTHIANS', 'ordinal_verse_start': 28778, 'chapters': 16}, {'book_name': '2CORINTHIANS', 'ordinal_verse_start': 29045, 'chapters': 13}, {'book_name': 'GALATIANS', 'ordinal_verse_start': 29190, 'chapters': 6}, {'book_name': 'EPHESIANS', 'ordinal_verse_start': 29339, 'chapters': 6}, {'book_name': 'PHILIPPIANS', 'ordinal_verse_start': 29444, 'chapters': 4}, {'book_name': 'COLOSSIANS', 'ordinal_verse_start': 29544, 'chapters': 4}, {'book_name': '1THESSALONIANS', 'ordinal_verse_start': 29623, 'chapters': 5}, {'book_name': '2THESSALONIANS', 'ordinal_verse_start': 29680, 'chapters': 3}, {'book_name': '1TIMOTHY', 'ordinal_verse_start': 29790, 'chapters': 6}, {'book_name': '2TIMOTHY', 'ordinal_verse_start': 29872, 'chapters': 4}, {'book_name': 'TITUS', 'ordinal_verse_start': 29925, 'chapters': 3}, {'book_name': 'PHILEMON', 'ordinal_verse_start': 29940, 'chapters': 1}, {'book_name': 'HEBREWS', 'ordinal_verse_start': 30243, 'chapters': 13}, {'book_name': 'JAMES', 'ordinal_verse_start': 30356, 'chapters': 5}, {'book_name': '1PETER', 'ordinal_verse_start': 30467, 'chapters': 5}, {'book_name': '2PETER', 'ordinal_verse_start': 30524, 'chapters': 3}, {'book_name': '1JOHN', 'ordinal_verse_start': 30626, 'chapters': 5}, {'book_name': '2JOHN', 'ordinal_verse_start': 30647, 'chapters': 1}, {'book_name': '3JOHN', 'ordinal_verse_start': 30660, 'chapters': 1}, {'book_name': 'JUDE', 'ordinal_verse_start': 30674, 'chapters': 1}, {'book_name': 'REVELATION', 'ordinal_verse_start': 31082, 'chapters': 22}]

const bookButton = document.getElementById('book-choose');
const chapterButton = document.getElementById('chapter-choose');
for(i=0; i<books.length; i++){
    a = document.createElement('option');
    a.textContent = books[i].book_name;
    bookButton.appendChild(a);
}
// its gonna start with Genesis so make chapter size 50
for(i=1; i<=50;i++){
    a = document.createElement('option');
    a.textContent = i
    chapterButton.appendChild(a);
}
// books.forEach(element => console.log(element));
const bookChoose = document.getElementById('book-choose');
const poemDisplay = document.getElementById('scripture');
const bookName = bookButton.value;
const chapterNum = chapterButton.value;
const searchHandler = document.getElementById('search-text');

console.log("bookname: ", bookName);
console.log("chaper: ", chapterNum);

bookChoose.addEventListener('change', () => {
    // clear all elements in scripture area
    document.getElementById('scripture').innerHTML = '';
    // set the verse value
    const placeholder = `https://mintz5.duckdns.org/bible/${bookButton.value}/${chapterButton.value}`;
    console.log(`url: ${placeholder}`);
    // updateDisplay(placeholder);
    const clearText = poemDisplay.querySelectorAll('p');
    clearText.forEach(function(item){
        item.remove();
    })
    // populate the text with Chapter 1 since its new book
    fetch(`https://mintz5.duckdns.org/bible/${bookButton.value}/1`)
        .then(response => {
            if(!response.ok){
                throw new Error(`HTTP Error: ${response.status}`);
            }
            return response.json();
        })
        .then(function (data){
            for(i=0; i<data.length;i++){
                const p = document.createElement('p');
                p.textContent = `${i+1}: ${data[i].text}`;
                poemDisplay.appendChild(p);
            }
        })
    // update the chapter selection
    for(i=0; i<books.length; i++){
        if(books[i].book_name == bookButton.value){
            let chapterMax = books[i].chapters
            chapterHandle = chapterButton.querySelectorAll('option');
            chapterHandle.forEach(function(item){
                item.remove();
            })
            for(j=1; j<=chapterMax; j++){
                co = document.createElement('option');
                co.textContent = j;
                chapterButton.appendChild(co);
            }
        }
    }
    // update the title to match book name
    document.getElementById('theTitle').textContent = bookButton.value;
});

chapterButton.addEventListener('change', () =>{
    // clear scripture area
    document.getElementById('scripture').innerHTML = '';
    const titleHandle = document.getElementById('theTitle');
    titleHandle.textContent = bookButton.value;
    const verse = bookChoose.value;
    const placeholder = `https://mintz5.duckdns.org/bible/${bookButton.value}/${chapterButton.value}`;
    console.log(verse);
    console.log(`url: ${placeholder}`);
    // updateDisplay(placeholder);
    const clearPoemDisplay = poemDisplay.querySelectorAll('p');
    clearPoemDisplay.forEach(function (item){
        item.remove();
    })
    fetch(placeholder)
        .then(response => {
            console.log("calling url");
            if(!response.ok){
                throw new Error(`HTTP error: ${response.status}`);
            }
            return response.json();
        })
        .then(function (data){
            for(i=0; i<data.length;i++){
                const p = document.createElement('p');
                p.textContent = `${i+1}: ${data[i].text}`;
                poemDisplay.appendChild(p);
            }
        })
})

function updateDisplay(verse){
    // clear the display area
    const titleHandle = document.getElementById('theTitle');
    titleHandle.textContent = bookButton.value;
    const clearPoemDisplay = poemDisplay.querySelectorAll('p');
    clearPoemDisplay.forEach(function (item){
        item.remove();
    })
    fetch(verse)
        .then(response => {
            console.log("calling url");
            if(!response.ok){
                throw new Error(`HTTP error: ${response.status}`);
            }
            return response.json();
        })
        .then(function (data){
            for(i=0; i<data.length;i++){
                const p = document.createElement('p');
                p.textContent = `${i+1}: ${data[i].text}`;
                poemDisplay.appendChild(p);
            }
        })
    // update chapter selection count
    clearChapters = chapterButton.querySelectorAll('option');
    clearChapters.forEach(function(item){
        item.remove();
    })
    // update chapters with respective book
    console.log("POPULATING CHAPTERS NOW");
    for(i=0; i<books.length;i++){
        console.log("looking at book ", books[i].book_name);
        if (books[i].book_name == bookButton.value){
            console.log(`found book: ${books[i].book_name}`)
            for(j=1; j<=books[i].chapters; j++){
                o = document.createElement('option');
                o.textContent= j;
                chapterButton.appendChild(o);
            }
            break;
        }
    }
} 
