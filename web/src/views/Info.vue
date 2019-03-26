<template>
  <div id="info">
    <talks />
  </div>
</template>

<script>
  export default {
    name: 'Info',

    components: {
      Talks: () => import('@/components/info/Talks')
    }
  }
</script>
