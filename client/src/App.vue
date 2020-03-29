<template>
  <div id="app">
    <div class="container">
      <div class="column">
        <ListWidget eventId="blogRoll" />
      </div>
      <div class="column y2">
        <Clock
          clockOneTz="America/Los_Angeles"
          clockThreeTz="America/New_York"
          eventId="clockWidget"
        />
      </div>
      <div class="column">
        <TextWidget />
      </div>
      <div class="column">
        <TextWidget title="Second" subtitle="Widget" />
      </div>
      <div class="column">
        <TextWidget title="Another" subtitle="Widget" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Clock from "./components/Clock.vue";
import TextWidget from "./components/TextWidget.vue";
import ListWidget from "./components/ListWidget.vue";
import { EventSink, eventType } from "./eventsink";
@Component({
  components: {
    Clock,
    TextWidget,
    ListWidget
  }
})
export default class App extends Vue {
  private eventServer = new EventSink();
  mounted() {
    (async () => {
      await this.eventServer.init((message: eventType<object>) =>
        this.$store.dispatch("change", message)
      );
    })();
  }
  beforeDestory() {
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
  height: 100vh;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 4px;
  width: 100%;
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

.center {
  display: flex;
  align-items: center;
  > div {
    text-align: center;
    width: 100%;
  }
}
</style>
