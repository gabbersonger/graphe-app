![Graphe](https://raw.githubusercontent.com/gabbersonger/graphe-app/main/build/banner.png)
<br />

# What is Graphe?

> Graphe is a **_fast and minimalistic Bible study application_**, with a focus on using the original language texts (the Hebrew Old Testament, the Septuagint and the Greek New Testament) in conjunction with an English translation. It is currently under active development.

Features include:

- Support for ESV, GNT and LXX
- Instant details while hovering original language
- Minimalistic design with built-in themes and fonts

<br/>

# Development & Building

Live development mode: `wails3 dev`

Build distributable app: `wails3 build`

Currently requires a large database file that is currently not packaged with the code. However there are plans to make this downloadable.

<br/>

# Roadmap

**Version 0.5 - Analysis**
- [x] Instant details timing
- [ ] Appearence
- [ ] Functionality of chapter data
  - [ ] Most popular words
  - [ ] Overlap with another chapter
  - [ ] Sort by: frequency (most -> least), frequency (least -> most), alphabetical
  - [ ] Filter out regular occurences
- [ ] Right click functionality

**Version 0.6 – Formatting Functions update**

- [ ] Functionality, e.g...
  - [ ] N occurances
  - [ ] Overlap with another chapter
  - [ ] Morph data
  - [ ] Specific word/root (incl. original languages)
- [ ] Combination of functions
- [ ] Modal
- [ ] Saved functions
- [ ] Settings (e.g. default colours)

<br/>

**Version 0.7 – Search update**

- [ ] Functionality, e.g...
  - [ ] Range
  - [ ] Items (e.g. topic, characters, root, strongs)
  - [ ] Commands (e.g. AND, OR, GROUPING, NOT, JOIN)
- [ ] Modal
- [ ] Settings (e.g. ranges, default text)

<br/>

**Version 0.8 - Settings update**
- [ ] General
  - [ ] Startup - what to show
  - [ ] Logs - open, purge
- [ ] Version / updates
  - [ ] Show version number
  - [ ] Check for updates at login
  - [ ] Application version
  - [ ] Database version
  - [ ] User login details
- [ ] Instant details
  - [x] Back to working
  - [ ] Fix throttling
  - [ ] UI - drag and drop available fields
  - [ ] Functionality
  - [ ] Reset
- [ ] Shortcuts
  - [x] Shortcuts reset
  - [ ] Shortcuts update menubar

<br/>

**Version 0.9 – Text details update**

- [ ] Parallel texts
- [ ] Bible headings
- [ ] Footnoted variant spelling/meaning
- [ ] Conjoin word highlighting
- [ ] Lexicon data
- [ ] Morph code expansions
- [ ] Cross References
- [ ] Song of Songs: "SHE", "HE", "OTHERS"
- [ ] Psalm book headings
- [ ] Divine name
- [ ] LXX Capitalisation
- [ ] Hebrew
  - [ ] RTL text support
  - [ ] Data
  - [ ] GetSections
  - [ ] GetWord
<br/>

**Version 0.10 - Make it work for others**

- [ ] Auto-download database files on start-up
  - [ ] Functionality
  - [ ] UI

<br/>

**Improvements**

- Zoom that works just for bible window
- Setting to reset zoom to value on startup
- Scroll to specific verse (that's not at start of paragraph)
- Add more fonts
- Add more themes
- "< Close" button in sidebar for settings at small screen sizes
- Animations (e.g. going into settings, like Discord)
- Sound effects
- On shortcut page, disable shortcuts

**Bugs**

- ESV psalm 119:1
- Creating shortcut with ' key (as this screws with mysql string)
- LXX go to Susanna 1 fails (Susanna 1:1 does not exist) -> use go function CreateFirstValidRef()
