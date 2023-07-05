import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, RouterProvider} from "react-router-dom"
import Home from './Rutas/Home'
import LoginRegister from './Rutas/LoginRegister'
import Reserve from './Rutas/Reserve'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/loginandregister",
    element: <LoginRegister />,
  },
  {
    path: "/reserve",
    element: <Reserve />,
  }
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)

