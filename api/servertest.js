const express = require('express');
const fs = require('fs');
const app = express();
const port = 3000;

let countries;
let categories;

try {
    const countriesData = fs.readFileSync('./data_json.json', 'utf8');
    const categoriesData = fs.readFileSync('./categories_json.json', 'utf8');
    
    countries = JSON.parse(countriesData);
    categories = JSON.parse(categoriesData);
} catch (error) {
    console.error('Error reading or parsing files:', error);
    process.exit(1); // Exit the process if there's an error
}

app.get('/getRandomCountry', (req, res) => {
    const randomCountry = getRandomElement(countries);

    console.log('Random Country:', randomCountry);

    res.json({ randomCountry });
});

app.get('/getRandomCategories', (req, res) => {
    const numberOfCategories = Math.floor(Math.random() * (10 - 5 + 1) + 5); // Random number between 5 and 10
    const randomCategories = getRandomElements(categories, numberOfCategories);

    console.log('Random Categories:', randomCategories);

    res.json({ randomCategories });
});

function getRandomElement(array) {
    const randomIndex = Math.floor(Math.random() * array.length);
    return array[randomIndex];
}

function getRandomElements(array, numberOfElements) {
    const shuffledArray = array.slice().sort(() => Math.random() - 0.5);
    return shuffledArray.slice(0, numberOfElements);
}

app.listen(port, () => {
    console.log(`Server is listening at http://localhost:${port}`);
});