const express = require('express');
const app = express();

// Mock users data
const users = {
  '1': { name: 'Alice', password: 'password1', email: 'alice@example.com' },
  '2': { name: 'Bob', password: 'password2', email: 'bob@example.com' },
};

app.get('/user/:id', (req, res) => {
  const user = users[req.params.id];
  if (user) {
    res.json(user);
  } else {
    res.status(404).send('User not found');
  }
});

app.listen(3000, () => console.log('Server running on port 3000'));
