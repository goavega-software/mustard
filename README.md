# Mustard
Mustard is an awesome dashboard inspired by Dashing written in Go and VueJS.

## Anatomy
Mustard primarily consists of 3 parts:
1. UI widgets - written in VueJS and TypeScript
2. Go Jobs - these are backend jobs which execute on a schedule to get the data
3. Server Sent Events is used to notify the UI widgets once new data is available on the server
## Config
* Mustard relies on dotenv for setting configuration values like API keys etc. There's a sample .env.example file which needs to be renamed as .env with all keys set.
* Mustard relies on cron to fire events on schedule but also exposes an API to fire jobs immediately (POST /api/nudge). This API is used by the Vue client to get the data on first run instead of waiting for the events to fire on schedule.
### First steps
```bash
$ cd <<folder>>/client
$ yarn install
$ yarn prod
$ cd ../
$ go mod download
$ go run main.go ./core ./jobs
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
### Creating new job
* Use any of the job as reference, move any configuration field to .env. You can read the env var using ``` os.Getenv("VAR") ```
* The job schedule can be defined in human readble form like "every nm|h|d"
* Call ```mustardcore.SseNotify(data)``` to notify the dashboard. eventId in data is what ties this data to a particular widget.
### Creating new widget
* Make a copy of TextWidget.vue and use that as reference
* Import the widget in App.vue and add it in the grid:
```html
      <div class="column">
        <WeatherWidget eventId="weather" />
      </div>
```
* Widgets use VueX for state management. Since all state objects have to be defined in advance, an empty defintion is requred in client/src/store/index.ts for the given event
* eventId prop ties this widget with SSE data
* widget width and height can be adjusted by setting classes x{1..4} and y{1..4} respectively.

Detailed wiki TBD!

### TODO
 - Create wiki
 - Drag and drop support
 - Better error handling - currently a job crash causes app to panic
 - Support for multiple dashboards
 - Data persistence
 - Support for widget installation via gists

### FAQ
Q. Why is it called Mustard?
A. Good question - we like using color names.
Q. Why Go and Vue, and not React and Node or X & Y?
A. This was primarily used as a weekend project to learn something new. Go and Vue seemed to be good choices to learn over the weekend :).
Q. How is it currently used?
A. Currently Mustard runs on desk on a repurposed screen from a bricked laptop and raspberry pi.

#### License
MIT