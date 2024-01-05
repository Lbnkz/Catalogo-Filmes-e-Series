import api from "./services";

export const handleLogin = async (username, password) => {
  try {
    // Faz uma requisição para obter todos os produtores da API
    const response = await api.get('/api/produtoras');
    
    console.log('API Response:', response.data);

    if (!response.data) {
      throw new Error('Erro ao obter produtores da API');
    }

    // Verifica se o produtor e senha correspondem a algum produtor na lista
    const produtorExists = response.data.some(produtor => {
      const match =
        produtor.Username.trim() === username.trim() &&
        produtor.Password.trim() === password.trim();
      return match;
    });
    return produtorExists;
  } catch (error) {
    console.error('Erro durante a autenticação:', error);
    return false;
  }
};