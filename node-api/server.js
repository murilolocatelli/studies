const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
const requireDir = require('require-dir');

// Iniciando o app
const app = express();

// Configuração para suportar requisições em JSON
app.use(express.json());
app.use(cors());

// Iniciando o DB
mongoose.connect(
  'mongodb://localhost:27017/nodeapi',
  {
    useNewUrlParser: true,
    useUnifiedTopology: true
  });

  requireDir('./src/model');

// Rotas
app.use('/api', require('./src/routes'));

app.listen(3000);
