<template>
  <div class="center" v-bind:style="containerStyle">
    <div>
      <h4>{{ model.subtitle }}</h4>
      <h2>{{ model.title }}</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { getStoreItem, State } from "../store";
type TextModel = { title: string; subtitle: string };
@Component
export default class TextWidget extends Vue {
  @Prop() private title?: string;
  @Prop() private subtitle?: string;
  @Prop() private eventId?: string;
  @Prop() private background?: string;
  get model(): TextModel {
    return (
      getStoreItem((this.$store.state as unknown) as State, this.eventId) || {
        title: this.title!,
        subtitle: this.subtitle!
      }
    );
  }

  get containerStyle() {
    return {
      background: this.background ? this.background : "#2a9d8f"
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
