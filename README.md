![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

# What is Graphe?

> Graphe is a **_fast and minimalistic Bible study application_**, with a focus on using the original language texts (the Hebrew Old Testament, the Septuagint and the Greek New Testament) in conjunction with an English translation. It is currently under active development.

Features include:

- Support for ESV, GNT and LXX
- Instant details while hovering original language
- Minimalistic design with built-in themes

<br/>

# Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

Test go code: `go test ./... -v`

Currently requires a large database file that is currently not packaged with the code. However there are plans to make this downloadable.

<br/>

# Roadmap

**Version 0.3 – Settings update**

- [x] Settings window
- [x] Saving settings between sessions
- [ ] Settings page
  - [x] Basic UI
  - [ ] General (startup)
  - [x] Appearence
    - [x] Scaling/zoom
    - [x] Font family
    - [x] Theme selector
  - [ ] Shortcuts
  - [ ] Formatting (e.g. default colours)
  - [ ] Search (e.g. ranges, default text)
  - [ ] Instant details (what is display and order?)
  - [ ] Version/updates
- [ ] Menubar
- [ ] Right click functionality

<br/>

**Verse 0.4 - Quality of life update**

- [ ] Remove JS Scripture functions
- [ ] Replace Virtualiser with just loading 10 chapters either side
- [ ] Animations
  - [ ] Entering settings window (look at what Discord does)
- [ ] Sound effects
- [ ] Auto-download database files on start-up
  - [ ] Functionality
  - [ ] UI

<br/>

**Version 0.5 – Text details update**

- [ ] LXX Capitalisation
- [ ] Parallel texts
- [ ] Bible headings
- [ ] Footnoted variant spelling/meaning
- [ ] Conjoin word highlighting
- [ ] Lexicon data
- [ ] Morph code expansions
- [ ] Cross References
- [ ] Analysis window in sidebar: top occurances in chapter (filter out regular occurances)
- [ ] Song of Songs: "SHE", "HE", "OTHERS"
- [ ] Psalm book headings
- [ ] Divine name

<br/>

**Version 0.6 – Functions update**

- [ ] Functionality, e.g...
  - [ ] N occurances
  - [ ] Overlap with another chapter
  - [ ] Morph data
  - [ ] Specific word/root (incl. original languages)
- [ ] Combination of functions
- [ ] Modal
- [ ] Saved functions

<br/>

**Version 0.7 – Search update**

- [ ] Functionality, e.g...
  - [ ] Range
  - [ ] Items (e.g. topic, characters, root, strongs)
  - [ ] Commands (e.g. AND, OR, GROUPING, NOT, JOIN)
- [ ] Modal

<br/>

**Version 0.8 - Hebrew update**

- [ ] RTL text support
- [ ] Tyndale HOT
  - [ ] Data
  - [ ] Displaying sections
  - [ ] Instant details

<br/>

**Later Versions**

- [ ] Multiple app states (sidebar keeps them all) - ⌘1, ⌘2, ⌘3 to cycle through screens

<br/><br/>

**Improvements**

- [ ] Zoom that works just for bible window


**Bugs**

- [ ] Resizing window does not keep verse at top (currently: only keeps block + scroll offset)
- [ ] Current verse doesn't update for book title
