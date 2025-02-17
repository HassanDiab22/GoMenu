import React, { useState } from "react";
import MenuCard from "./MenuCard"; // Assuming MenuCard is in the same directory

const exampleMenus = [
  {
    Url: "https://via.placeholder.com/150",
    Name: "Burger",
  },
  {
    Url: "https://via.placeholder.com/150",
    Name: "Pizza",
  },
  {
    Url: "https://via.placeholder.com/150",
    Name: "Pasta",
  },
];

export default function MenusPage() {
  const [menus, setMenus] = useState(exampleMenus);

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="w-full px-6 py-8 md:h-screen lg:py-0">
        <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
              {menus.map((menu, index) => (
                <MenuCard key={index} menu={menu} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
