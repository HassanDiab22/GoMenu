/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./resources/**/*.blade.php",
    "./resources/**/*.js",
    "./resources/**/*.vue",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#454", // Custom primary color (blue-900)
        secondary: "#9333EA", // Custom secondary color (purple-600)
      },
    },
  },
  plugins: [],
};
