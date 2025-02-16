import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { publicRoutes } from "./publicRoutes";
export default function RouteConfig() {
  return (
    <Router>
      <Routes>
        {publicRoutes.map((route, index) => (
          <Route key={index} path={route.path} element={<route.component />} />
        ))}
      </Routes>
    </Router>
  );
}
