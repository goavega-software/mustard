<template>
  <h1>
    <span>{{ model.up ? "&#9650;" : "&#9660;" }}</span>
    {{ model.tweeningValue }}
  </h1>
</template>
<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import { scriptLoader } from "../../script-loader";
declare global {
  interface Window {
    TWEEN: {
      Tween: (options: any) => void;
      update: (current: any) => boolean;
    };
  }
}
type numberModel = { up: boolean; tweeningValue: number };
@Component
export default class AnimatedNumberComponent extends Vue {
  @Prop() private value?: number;
  private model: numberModel = {
    up: true,
    tweeningValue: 0
  };

  mounted() {
    (async () => {
      await scriptLoader("https://cdn.jsdelivr.net/npm/tween.js", true);
      this.tween(0, this.value || 0);
    })();
  }
  @Watch("value")
  onPropertyChanged(newValue: number, oldValue: number) {
    this.model.up = newValue >= oldValue;
    this.tween(oldValue, newValue);
  }
  tween(startValue: number, endValue: number) {
    let animationFrame: number;
    const animate = (currentTime: number) => {
      if (window.TWEEN.update(currentTime)) {
        animationFrame = requestAnimationFrame(animate);
      }
    };
    // @ts-ignore
    new window.TWEEN.Tween({ tweeningValue: startValue })
      .to({ tweeningValue: endValue }, 500)
      .onUpdate((value: number) => {
        const tweeningValue = +(value * endValue).toFixed(0);
        this.model = { ...this.model, tweeningValue };
      })
      .onComplete(() => cancelAnimationFrame(animationFrame))
      .start();
    animationFrame = requestAnimationFrame(animate);
  }
}
</script>
