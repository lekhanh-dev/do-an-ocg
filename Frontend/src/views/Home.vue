<template>
  <div class="page">
    <SlideImage />
    <ProductList />
  </div>
</template>

<script>
import SlideImage from "@/components/SlideImage.vue";
import ProductList from "@/components/ProductList.vue";

export default {
  name: "Home",
  components: {
    SlideImage,
    ProductList,
  },
};
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
}
</style>
