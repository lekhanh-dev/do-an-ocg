import { createStore } from "vuex";

export default createStore({
  state: {
    cart: [],
  },
  getters: {
    getSumProduct: (state) => {
      return state.cart.reduce((acc, product) => acc + product.num, 0);
    },
    getTotal: (state) => {
      return state.cart.reduce(
        (acc, product) => acc + product.num * product.Price,
        0
      );
    },
  },
  mutations: {
    addToCart(state, { product, num }) {
      const productObj = state.cart.find((p) => p.ID === product.ID);
      if (productObj) {
        productObj.num += num;
      } else {
        state.cart.unshift({ ...product, num: num });
      }
    },
    updateNumProduct(state, { id, num }) {
      const product = state.cart.find((p) => p.ID === id);
      product.num = num;
    },
    deleteProduct(state, id) {
      state.cart = state.cart.filter((product) => product.ID !== id);
    },
  },
  actions: {},
  modules: {},
});
