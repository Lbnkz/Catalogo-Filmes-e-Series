import axios from "axios";

/**
 * @function api
 * 
 * @brief Cria uma instância do AXIOS, gerando conexão com a API.
*/
const api = axios.create({
  baseURL: "http://localhost:8085"
});



export default api;
