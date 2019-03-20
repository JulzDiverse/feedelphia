<template>
  <div id="home">
    <articles>
      <banner />
    </articles>
  </div>
</template>

<script>
  export default {
    name: 'Home',

    components: {
      Articles: () => import('@/components/home/Articles'),
      Banner: () => import('@/components/home/Banner')
    }
  }
</script>
