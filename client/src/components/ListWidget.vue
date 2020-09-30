<template>
  <div class="blogroll">
    <h3>{{ model.title }}</h3>
    <transition name="slide-fade" mode="out-in">
      <div :key="index">
        <h4 v-if="item != undefined">{{ item.title }}</h4>
        <img
          v-if="item != undefined"
          :src="item.url"
          style="float:left;margin-right:5px"
        />
        <p v-if="item != undefined">{{ item.description | subString }}</p>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { getStoreItem, State } from "../store";
type listItem = { title: string; description: string; url: string };
@Component({
  filters: {
    subString(s: string) {
      return `${s.substr(0, 180)}...`;
    }
  }
})
export default class ListWidget extends Vue {
  @Prop() eventId?: string;
  private item: listItem = { title: "", description: "", url: "" };
  private timerId?: number;
  private index = -1;
  get model(): { title: string; items: [] } {
    if (this.timerId) {
      clearInterval(this.timerId);
    }
    this.timerId = setInterval(this.cycle, 10000);
    return (
      getStoreItem((this.$store.state as unknown) as State, this.eventId) || {
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
    const item: listItem = this.model.items[this.index];
    if (!item) {
      return;
    }
    this.item = { ...item };
    // Vue.set(this.item, "title", item.title);
    // Vue.set(this.item, "description", item.description);
    // Vue.set(this.item, "url", item.url);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.blogroll {
  background: #e76f51;
  padding: 10px;
  h3 {
    text-align: center;
  }
  h4 {
    padding: 10px;
  }
  p {
    text-justify: distribute;
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
