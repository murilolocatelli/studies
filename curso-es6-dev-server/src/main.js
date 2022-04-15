const minhaPromisse = () => 
  new Promise((resolve, reject) => setTimeout(() => resolve('ok'), 2000));

minhaPromisse()
  .then(response => {
    
    console.log(response);

    minhaPromisse().then(response => {

      console.log(response);

      minhaPromisse().then(response => {

        console.log(response)

        executarPromisse();
      });
    });
  });

/*async function executarPromisse() {
  console.log(await minhaPromisse());
  console.log(await minhaPromisse());
  console.log(await minhaPromisse());
}*/

// Mesmo comportamento que a sintaxe acima
const executarPromisse = async () => {
  console.log(await minhaPromisse());
  console.log(await minhaPromisse());
  console.log(await minhaPromisse());
}

//Erro, pois não está dentro de uma função com async

//console.log(await minhaPromisse());

import axios from 'axios'

class Api {
  static async getUserInfo(username) {
    
    try {
      const response = await axios.get(`https://api.github.com/users/${username}`);

      console.log(response);

    } catch (err) {
      console.warn(err);
    }
  }
}

Api.getUserInfo('murilolocatellifasdf');