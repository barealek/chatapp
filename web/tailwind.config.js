import containerQueries from "@tailwindcss/container-queries";
import typography from "@tailwindcss/typography";
import forms from '@tailwindcss/forms';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],

  theme: {
    extend: {}
  },

  plugins: [forms, typography, containerQueries]
};
