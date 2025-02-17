import Signin from "../app/Auth/Signin";
import MenusPage from "../app/Menus/MenusPage";

export const publicRoutes = [
  { path: "/", component: Signin },
  { path: "/menus", component: MenusPage },
];
