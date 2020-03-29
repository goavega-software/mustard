<template>
  <div class="blogroll">
    <h3>{{ model.title }}</h3>
    <transition name="slide-fade" mode="out-in">
      <div :key="index">
        <h4 v-if="item != undefined">{{ item.title }}</h4>
        <p v-if="item != undefined">{{ item.description }}</p>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
type listItem = { title: string; description: string; link: string };
@Component
export default class ListWidget extends Vue {
  @Prop() eventId?: string;
  private model: { title: string; items: [] };
  private item?: listItem = { title: "", description: "" };
  private timerId?: number;
  private index = -1;
  get model() {
    if (this.timerId) {
      clearInterval(this.timerId);
    }
    this.timerId = setInterval(this.cycle, 10000);
    return (
      this.$store.state[this.eventId || "undefined"] || {
        title: "Hello",
        items: []
      }
    );
  }
  cycle() {
    if (!this.model.items || this.model.items.length === 0) {
      return;
    }
    this.index++;
    if (this.index > this.model.items.length) {
      this.index = 0;
    }
    Vue.set(this.item, "title", this.model.items[this.index].title);
    Vue.set(this.item, "description", this.model.items[this.index].description);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.blogroll {
  background: #2a9d8f;
  h3 {
    text-align: center;
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
