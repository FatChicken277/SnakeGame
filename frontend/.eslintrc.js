module.exports = {
  env: {
    es6: true,
    node: true,
  },
  extends: [
    '@vue/airbnb',
    'plugin:vue/essential',
  ],
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
  },
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: 'module',
  },
  plugins: [
    'vue',
  ],
  rules: {
  },
};
