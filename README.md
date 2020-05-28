# Mustard
Mustard is an awesome dashboard written in Go, TypeScript and VueJS. Mustard uses [CSS Grid layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout) hence only supports modern browsers.

![screenshot](https://github.com/goavega-software/mustard/raw/master/client/public/screenshot.png)

## Anatomy
Mustard has 4 components:
1. UI widgets - written in VueJS and TypeScript
2. Go Jobs - these are backend jobs which execute on a schedule to get the data
3. A Kafka topic consumer
4. Server Sent Events for pushing data to the UI widgets
## Config
* Mustard relies on dotenv for setting configuration values like API keys etc. There's a sample .env.example file which needs to be renamed as .env with all keys set.
* Mustard relies on cron to fire events on schedule but also exposes an API to fire jobs immediately (POST /api/nudge). This API is used by the Vue client to get the data on first run instead of waiting for the events to fire on schedule.
* Since v0.2 Mustard also supports consuming Kafka topics and forwarding those events over Server Sent Events. More details are under Kafka section.

### First steps
```bash
$ cd <<folder>>/client
$ yarn install
$ yarn prod
$ cd ../
$ go mod download
$ go build
$ ./mustard
```
Navigate to http://localhost:8090/dashboard. Default port is 8090 unless overridden. 
## Quick start
Below is how the folder structure looks like:
```bash
├───client
│   ├───node_modules
│   └───src
│       ├───assets
│       ├───components
│       ├───eventsink
│       └───store
└───jobs
```
Folders of interest are:
1. client/src/components - All dashboard widgets should be in this folder.
2. jobs/ - All Go jobs should be in this folder.
### OOBE
Few widgets and jobs are already included in repository - feel free to use/abuse them 
#### Widgets
* Text widget - displays a title and subtitle
* Slideshow widget - cycles through images on an interval. This is currently wired up with a job to get flickr images 
* Clock widget - displays 3 clocks and a trivia on today's day
* A comparer widget - Displays a stacked bar graph comparing 2 values. If the difference between the values is > pre-defined threshold then entire widget turns red else stays green. Currently it display the # of Covid-19 cases in india for T+0 and T-1.
* Weather widget - displays weather from OpenWeather and gets the backrground image from flickr for the weather condition
* List widget - cycles through a list of items (title, image and description) on schedule
#### Jobs
* Blogroll - gets RSS feed from a site and converts the URL into a QR Code.
* Weather - gets the current weather from OpenWeather
* Number - gets the number trivia for today's day
* comparison - gets today and yesterday's Covid-19 cases
* FlickrShow - gets images from flickr
* S3Show - gets images from a S3 bucket filtered by a prefix (aka folder).
### Creating new job
* Use any of the job as reference, move any configuration field to .env. You can read the env var using ``` os.Getenv("VAR") ```
* The job schedule can be defined in human readble form like "every nm|h|d"
* Call ```mustardcore.GetEventsManager().Notify(data)``` to notify the dashboard. eventId in data is what ties this data to a particular widget. Below is an example of how number job pushes data to clockWidget
```golang
data := mustardcore.EventData{Event: "clockWidget",  Data: number{Trivia: string(text)}}
mustardcore.GetEventsManager().Notify(data)
```
### Creating new widget
* Make a copy of TextWidget.vue and use that as reference
* Import the widget in App.vue and add it in the grid:
```html
      <div class="column">
        <WeatherWidget eventId="weather" />
      </div>
```
* Widgets use VueX for state management. Since all state objects have to be defined in advance, an empty defintion is required in client/src/store/index.ts for the given event.
* eventId prop ties this widget with SSE data
* widget width and height can be adjusted by setting classes x{1..4} and y{1..4} respectively.
* Starting v0.2, a shared animated number widget component is added. This component can be used to animate a changing number value. The widget is same as one provided in Vue state animation [example](https://vuejs.org/v2/guide/transitioning-state.html).

## Extending
Apart from the regular way of cloning and modifying, mustard supports extensions via gists. The gists can be installed by providing the ID of the gist ` $ ./mustard -gist={id} `. Only .vue and .go files are supported in the gist. Sample gists

| widget        | Description  |
| ------------- | -----|
| [Mopidy Client](https://gist.github.com/sn123/dc903fb6d19ad793505fecc1b52536f7)     |A Simple client for [Mopidy](https://mopidy.com/) which displays now playing|
| [Animated Number](https://gist.github.com/sn123/ad1b185cdd76abab4d0bd5eb4c8c4b22)     |Animates a number and displays up/down arrow|

### Kafka Support
* From v0.2, Mustard starts a kafka listener by connecting to broker URL specified in .env and subscribing to topic in .env. Multiple Broker URLs in .env should be separated by comma. Mustard kafka listener is primarily a proxy between the widget and the kafka topic, it does not add any intelligence on top of the message received and just forwards it over the SSE. Message published on kafka topic should have same JSON structure as expected by SSE:
```golang
type EventData struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
```
Kafka topics are great for pushing NRT metrices (Near Real Time) to the dashboard. Kafka topic bound widgets can be used to display Time series graphs or volatile values. Mustard already has built-in support for [ApexCharts](https://apexcharts.com/), which can be used for displaying charts.

### Docker compose
There's a docker compose file which exposes kafka listener both internally and externally. On windows, since traffic cannot be routed to linux containers, the kafka listener is exposed as host.docker.internal on the host. The docker-compose for windows is docker-compose-win.yml.

```shell
$ docker-compose -f ./docker-compose.yml up -d
```
Please make sure that mustard's env file has the topic you need to listen to (KAFKA_TOPIC)

The docker compose file uses https://github.com/wurstmeister/kafka-docker/, please refer to the documentation there to troubleshoot connectivity issues.

On windows, this is how it would potentially work:

1. Create the topic:
```.\kafka-topics.bat --bootstrap-server host.docker.internal:9094 --topic test --create```
2. Set the topic name in .env
3. Produce a message:
```.\kafka-console-producer.bat --bootstrap-server host.docker.internal:9094 --topic test```

### Docker
```sh
$ docker pull goavega/mustard:latest
$ docker run -p <local>:80 --env-file ./.env goavega/mustard
```

### TODO
 - Create wiki
 - Drag and drop support
 - Better error handling - currently a job crash causes app to panic
 - Support for multiple dashboards
 - Data persistence
 - Shared Chart components

### FAQ
Q. Why is it called Mustard?

A. Good question - we like using color names.

Q. Why Go and Vue, and not React and Node or X & Y?

A. This was primarily used as a weekend project to learn something new. Go and Vue seemed to be good choices to learn over the weekend :).

Q. How is it currently used?

A. Currently Mustard runs on desk on a repurposed screen from a bricked laptop and raspberry pi.


#### License
MIT. Inspired by Dashing.
