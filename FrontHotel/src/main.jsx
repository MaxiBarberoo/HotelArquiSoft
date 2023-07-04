import React from 'react'
import ReactDOM from 'react-dom/client'
import './Stylesheet/index.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom"
import Home from './Rutas/Home'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },

    {

      path: "/register",


    },

  {

    path: "/login"




  },

  {

    path: "/reserve"




  }

])



ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
      <RouterProvider router={router} />

  </React.StrictMode>,
)

