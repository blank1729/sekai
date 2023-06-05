import React from "react";
import ReactDOM from "react-dom/client";
import "./index.scss";
import "./reset.scss";
import { RouterProvider } from "react-router-dom";
import { router } from "./routes/routes.tsx";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
