const express = require('express');
const mysql = require('mysql');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());

const db = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'password',
    database: 'blogdb'
});

db.connect((err) => {
    if (err) throw err;
    console.log('Connected to database');
});

app.get('/api/search', (req, res) => {
    const query = req.query.q;
    db.query(`SELECT * FROM blogs WHERE title LIKE '%${query}%'`, (error, results) => {
        if (error) {
            console.log(error);
        } else {
            res.send(results);
        }
    });
});

app.listen(3000, () => {
    console.log('Server started on port 3000');
});
