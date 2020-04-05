<template>
  <div class="center blue">
    <div>
      <h4>{{ model.subtitle }}</h4>
      <h2>{{ model.title }}</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
type TextModel = { title: string; subtitle: string };
@Component
export default class TextWidget extends Vue {
  @Prop() private title?: string;
  @Prop() private subtitle?: string;
  @Prop() private eventId?: string;
  private model: TextModel;

  get model() {
    return (
      this.$store.state[this.eventId || "undefined"] || {
        title: this.title,
        subtitle: this.subtitle
      }
    );
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.blue {
  background: #2481b7;
}
</style>
