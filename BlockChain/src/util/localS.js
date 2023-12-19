const localS = {
  getItem: (key) => {
    return localStorage.getItem(key);
  },
  setItem: (key, value) => {
    return localStorage.setItem(key, value);
  },
  removeItem: (key) => {
    return localStorage.removeItem(key);
  },
};

export default localS;
