import React from "react";
import { Card, CardContent } from "@mui/material";

type Menu = {
  Url: string;
  Name: string;
};

interface MenuCardProps {
  menu: Menu;
}

const MenuCard: React.FC<MenuCardProps> = ({ menu }) => {
  return (
    <Card className="flex items-center p-4 shadow-lg rounded-lg">
      {/* Image */}
      <img
        src={menu.Url}
        alt={menu.Name}
        className="w-24 h-24 object-cover rounded-lg mx-auto"
      />

      {/* Name */}
      <CardContent className="ml-4">
        <h2 className="text-lg font-semibold">{menu.Name}</h2>
      </CardContent>
    </Card>
  );
};

export default MenuCard;
