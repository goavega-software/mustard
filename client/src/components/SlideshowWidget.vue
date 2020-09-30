<template>
  <div class="slideshow">
    <transition name="slide-fade" mode="out-in">
      <div :key="index" class="full">
        <div class="full" v-bind:style="containerStyle"></div>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { getStoreItem, State } from "../store";
type Slide = { url: string };
type SlideShowModel = { items: Array<Slide> };
@Component
export default class SlideshowWidget extends Vue {
  @Prop() eventId?: string;
  private item?: Slide = { url: "https://source.unsplash.com/random" };
  private timerId?: number;
  private index = -1;
  mounted() {
    if (this.timerId) {
      clearInterval(this.timerId);
    }
    this.timerId = setInterval(this.cycle, 10000);
  }
  get model(): SlideShowModel {
    return (
      getStoreItem((this.$store.state as unknown) as State, this.eventId) || {
        items: []
      }
    );
  }
  get containerStyle() {
    return {
      background:
        this.item && this.item.url
          ? `url(${this.item.url}) no-repeat`
          : "#1446a0"
    };
  }
  cycle() {
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
    this.item = { ...item };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.slideshow {
  background: linear-gradient(to right, #f3904f, #3b4371);
  overflow: hidden;
  .full {
    width: 100%;
    height: 100%;
    background-size: cover !important;
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
  transform: translateX(5px);
  opacity: 0;
}
</style>
