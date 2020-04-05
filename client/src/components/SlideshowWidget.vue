<template>
  <div class="slideshow">
    <transition name="slide-fade" mode="out-in">
      <div :key="index">
        <img v-if="item != undefined" :src="item.url" />
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
type Slide = { url: string };
type SlideShowModel = { items: Array<Slide> };
@Component
export default class SlideshowWidget extends Vue {
  @Prop() eventId?: string;
  private item?: Slide = { url: '' };
  private timerId?: number;
  private index = -1;
  get model(): SlideShowModel {
    if (this.timerId) {
      clearInterval(this.timerId);
    }
    this.timerId = setInterval(this.cycle, 1000);
    return (
      this.$store.state[this.eventId || 'undefined'] || {
        items: []
      }
    );
  }
  cycle() {
    console.log('model is', this.model);
    if (!this.model || !this.model.items || this.model.items.length === 0) {
      return;
    }
    this.index++;
    if (this.index > this.model.items.length) {
      this.index = 0;
    }
    const item = this.model.items[this.index];
    if (!item) {
      return;
    }
    console.log('item is', item);
    Vue.set(this.item, 'url', item.url);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.slideshow {
  background: linear-gradient(to right, #f3904f, #3b4371);
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active for <2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
</style>
