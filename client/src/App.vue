<template>
  <!-- <div id="app">
    <div class="container"> -->
      <grid-layout :layout.sync="layout"
              :col-num="12"
              :row-height="30"
              :is-draggable="draggable"
              :is-resizable="resizable"
              :vertical-compact="false"
              :use-css-transforms="false"
      >
        <grid-item v-for="(item, index) in layout"
                   v-bind:key=index
                   :static="item.static"
                   :x="item.x"
                   :y="item.y"
                   :w="item.w"
                   :h="item.h"
                   :i="item.i"
        >
          <span class="text">{{item.i}}</span>
        </grid-item>
    </grid-layout>
      <!-- <div
        v-for="(l, index) in layout"
        v-bind:key="index"
        v-bind:class="getClass(index)"
      >
        <div :is="l.component" v-bind="getProps(index)" />
      </div> -->

    <!-- </div>
  </div> -->
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
import { dashboardStaticConfig } from "./dashboard-static-config";
import { GridLayout, GridItem } from "vue-grid-layout";

type layoutType = {
  x: number;
  y: number;
  w: number;
  h?: number;
  i: string;
};
interface Classes {
  column: boolean;
  [key: string]: boolean;
}

@Component({
  components: {
    Clock,
    TextWidget,
    ListWidget,
    ComparisonWidget,
    WeatherWidget,
    SlideshowWidget,
    GridLayout,
    GridItem
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
      // this.layout = layout;

      console.log("🚀 ~ file: App.vue:79 ~ App ~ this.layout:", this.layout);
      // set the Vuex state modules
      // this.layout.forEach(layoutItem => {
      //   if (!layoutItem.state) {
      //     return;
      //   }
      //   const moduleState = {
      //     state: () => ({ data: {} })
      //   };
      //   const stateKey = Object.keys(layoutItem.state)[0];
      //   this.$store.registerModule(stateKey, moduleState);
      // });
    })();
    (async () => {
      await this.eventServer.init((message: eventType<object>) =>
        this.$store.dispatch("change", message)
      );
    })();
  }
  data() {
    return {
        layout: [
            {"x":0,"y":0,"w":2,"h":2,"i":"bsGenerator", static: false},
            {"x":2,"y":0,"w":2,"h":4,"i":"jotd", static: false},
            {"x":4,"y":0,"w":2,"h":5,"i":"clock", static: false}
        ],
        draggable: true,
        resizable: true,
        index: 0
    }
  }
  // getClass(index: number): { column: boolean } & Record<string, boolean> {
  //   const classes: Classes = {
  //     column: false
  //   };
  //   const classNames = Array.isArray(this.layout[index].class)
  //     ? this.layout[index].class
  //     : [this.layout[index].class];
  //   (classNames as string[]).forEach(cls => {
  //     classes[cls] = false;
  //   });
  //   return classes;
  // }
  // getProps(index: number): Record<string, unknown> {
  //   return this.layout[index].props;
  // }
  beforeDestroy() {
    this.eventServer.destory();
  }
}
</script>

<style lang="scss">
// body {
//   margin: 0;
//   padding: 0;
//   background: #333;
//   color: #efefef;
//   font-family: Karla, sans-serif;
// }

// body * {
//   margin: 0;
//   padding: 0;
//   box-sizing: border-box;
//   font-size: 15px;
// }

// $base: 20;

// @for $i from 1 through 4 {
//   h#{$i} {
//     font-size: $base + (4-$i) * 10px;
//   }
// }

// .container {
//   height: 100vh;
//   display: grid;
//   grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
//   gap: 4px;
//   width: 100%;
//   grid-auto-rows: 1fr;
// }

// .column {
//   @for $i from 1 through 4 {
//     &.x#{$i} {
//       grid-column-end: span $i;
//     }
//     &.y#{$i} {
//       grid-row-end: span $i;
//     }
//   }
//   & > div {
//     height: 100%;
//   }
// }
// .text-center {
//   text-align: center;
// }
// .center {
//   display: flex;
//   align-items: center;
//   > div {
//     text-align: center;
//     width: 100%;
//   }
// }
.vue-grid-layout {
    background: #eee;
}
.vue-grid-item:not(.vue-grid-placeholder) {
    background: #ccc;
    border: 1px solid black;
}
.vue-grid-item .resizing {
    opacity: 0.9;
}
.vue-grid-item .static {
    background: #cce;
}
.vue-grid-item .text {
    font-size: 24px;
    text-align: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    height: 100%;
    width: 100%;
}
.vue-grid-item .no-drag {
    height: 100%;
    width: 100%;
}
.vue-grid-item .minMax {
    font-size: 12px;
}
.vue-grid-item .add {
    cursor: pointer;
}
.vue-draggable-handle {
    position: absolute;
    width: 20px;
    height: 20px;
    top: 0;
    left: 0;
    background: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='10' height='10'><circle cx='5' cy='5' r='5' fill='#999999'/></svg>") no-repeat;
    background-position: bottom right;
    padding: 0 8px 8px 0;
    background-repeat: no-repeat;
    background-origin: content-box;
    box-sizing: border-box;
    cursor: pointer;
}
</style>
