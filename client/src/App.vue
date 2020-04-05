<template>
  <div id="app">
    <div class="container">
      <div class="column y2">
        <WeatherWidget eventId="weather" />
      </div>
      <div class="column y2">
        <Clock
          clockOneTz="America/Los_Angeles"
          clockThreeTz="America/New_York"
          eventId="clockWidget"
        />
      </div>
      <div class="column y1">
        <ListWidget eventId="blogRoll" />
      </div>
      <div class="column y1">
        <ComparisonWidget
          eventId="comparisonWidget"
          threshold="0.25"
          new="0"
          old="0"
        />
      </div>
      <div class="column x2 y1">
        <SlideshowWidget eventId="slideshow" />
      </div>
      <div class="column">
        <TextWidget title="universe" subtitle="hello" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Clock from './components/Clock.vue';
import TextWidget from './components/TextWidget.vue';
import ListWidget from './components/ListWidget.vue';
import ComparisonWidget from './components/ComparisonWidget.vue';
import WeatherWidget from './components/WeatherWidget.vue';
import SlideshowWidget from './components/SlideshowWidget.vue';
import { EventSink, eventType } from './eventsink';
import { BaseUrl } from './constants';
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
  created() {
    setTimeout(
      async () =>
        await fetch(`${BaseUrl}api/nudge`, {
          method: 'POST'
        }),
      1000
    );
  }
  mounted() {
    (async () => {
      await this.eventServer.init((message: eventType<object>) =>
        this.$store.dispatch('change', message)
      );
    })();
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
  height: 100vh;
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
