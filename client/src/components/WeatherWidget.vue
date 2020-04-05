<template>
  <div class="center yale" v-bind:style="containerStyle">
    <div>
      <img :src="model.icon" v-if="model.icon" />
      <h4>{{ model.desc }}</h4>
      <h2>{{ model.temp }} &#176;</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
type WeatherModel = {
  temp: number;
  desc: string;
  icon?: string;
  image?: string;
};
@Component
export default class TextWidget extends Vue {
  @Prop() private eventId?: string;
  get model(): WeatherModel {
    return (
      this.$store.state[this.eventId || 'undefined'] || {
        temp: 0,
        desc: '--'
      }
    );
  }

  get containerStyle() {
    return {
      background: this.model.image
        ? `url(${this.model.image}) no-repeat`
        : '#1446a0'
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.yale {
  background: #1446a0;
  background-size: cover !important;
  > div {
    background: rgba(0, 0, 0, 0.2);
    border-radius: 5px;
  }
}
</style>
