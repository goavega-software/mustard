<template>
  <div id="app">
    <div class="container">
      <div
        v-for="(l, index) in layout"
        v-bind:key="index"
        v-bind:class="getClass(index)"
      >
        <div :is="l.component" v-bind="getProps(index)" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Clock from "./components/Clock.vue";
import TextWidget from "./components/TextWidget.vue";
import ListWidget from "./components/ListWidget.vue";
import ComparisonWidget from "./components/ComparisonWidget.vue";
import WeatherWidget from "./components/WeatherWidget.vue";
import SlideshowWidget from "./components/SlideshowWidget.vue";
import { EventSink, eventType } from "./eventsink";
import { BaseUrl } from "./constants";
type layoutType = {
  component: string;
  class: string | string[];
  props: Record<string, unknown>;
  state?: Record<string, unknown>;
};
interface Classes {
  column: boolean;
  [key: string]: boolean;
}
const getLayout = async () => {
  const url = window.location.pathname;
  const path = url.substring(url.lastIndexOf("/") +1);
  const response = await fetch(`${BaseUrl}api/layout`, {
    method: "POST",
    headers: {
      "Content-type": "application/json"
    },
    body: JSON.stringify({ path })
  });
  return response.json() as Promise<layoutType[]>;
};
@Component({
  components: {
    Clock,
    TextWidget,
    ListWidget,
    ComparisonWidget,
    WeatherWidget,
    SlideshowWidget
  }
})
export default class App extends Vue {
  private eventServer = new EventSink();
  private layout: Array<layoutType> = [];

  created() {
    setTimeout(
      async () =>
        await fetch(`${BaseUrl}api/nudge`, {
          method: "POST"
        }),
      1000
    );
  }
  mounted() {
    (async () => {
      // get the layout, for now we get it from local
      this.layout = await getLayout();
      // set the Vuex state modules
      this.layout.forEach(layoutItem => {
        if (!layoutItem.state) {
          return;
        }
        const moduleState = {
          state: () => ({ data: {} })
        };
        const stateKey = Object.keys(layoutItem.state)[0];
        this.$store.registerModule(stateKey, moduleState);
      });
    })();
    (async () => {
      await this.eventServer.init((message: eventType<object>) =>
        this.$store.dispatch("change", message)
      );
    })();
  }
  getClass(index: number): { column: boolean } & Record<string, boolean> {
    const classes: Classes = {
      column: true
    };
    const classNames = Array.isArray(this.layout[index].class)
      ? this.layout[index].class
      : [this.layout[index].class];
    (classNames as string[]).forEach(cls => {
      classes[cls] = true;
    });
    return classes;
  }
  getProps(index: number): Record<string, unknown> {
    return this.layout[index].props;
  }
  beforeDestroy() {
    this.eventServer.destory();
  }
}
</script>

<style lang="scss">
body {
  margin: 0;
  padding: 0;
  background: #333;
  color: #efefef;
  font-family: Karla, sans-serif;
}

body * {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-size: 15px;
}

$base: 20;

@for $i from 1 through 4 {
  h#{$i} {
    font-size: $base + (4-$i) * 10px;
  }
}

.container {
  height: 100vh;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 4px;
  width: 100%;
  grid-auto-rows: 1fr;
}

.column {
  @for $i from 1 through 4 {
    &.x#{$i} {
      grid-column-end: span $i;
    }
    &.y#{$i} {
      grid-row-end: span $i;
    }
  }
  & > div {
    height: 100%;
  }
}
.text-center {
  text-align: center;
}
.center {
  display: flex;
  align-items: center;
  > div {
    text-align: center;
    width: 100%;
  }
}
</style>
