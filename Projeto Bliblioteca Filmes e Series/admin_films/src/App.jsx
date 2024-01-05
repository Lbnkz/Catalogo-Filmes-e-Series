// Import de Bibliotecas
import React from "react";
import { Route, Routes } from "react-router-dom";

// Import de Componentes
import Login from "./pages/Login/Login";
import UserScreen from "./pages/UserScreen/UserScreen";
import NoMatch from "./pages/NoMatch/NoMatch";

import { Toaster } from "react-hot-toast";
import Usos from "./pages/catalogo/usos";
import CadastroProdutora from "./pages/UserScreen/CadastroProdutora";
import EditProdutora from "./pages/UserScreen/EditProdutora";

/**
 * @brief Corpo principal da aplicaçõa React. Através deste elemento, são roteadas as telas de acordo com a URL inserida pelo usuário no navegador.
 * @extends react-router-dom (V6) Utiliza da biblioteca react-router-dom na versão 6.x.x para poder gerenciar o roteamento, através da tag "Routes".
 */
function App() {
  return (
    <div className="App">
      <Toaster />
      <Routes>
        <Route path="/" element={<Login />} />

        <Route
          default
          path="produtoras"
          element={
            <UserScreen />
          }
        />

        <Route
          path="produtoras/cadastro"
          element={
            <CadastroProdutora />
          }
        />

        <Route
          path="produtoras/edit/:id"
          element={
            <EditProdutora />
          }
        />
        <Route
          path="catalogo"
          element={
            <Usos />
          }
        />
        <Route path="*" element={<NoMatch />} />
      </Routes>
    </div>
  );
}

export default App;
