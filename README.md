![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

## Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

## Roadmap

**Version 1 Checklist**

- [x] Add keyboard shortcuts
- [x] Format display of texts
  - [x] Update virtualiser
  - [x] Add book titles
  - [x] Add top padding
- [x] Get text/passage selection working
  - [x] Functionality
  - [x] 'Choose version' modal
  - [x] 'Choose text' modal
- [x] Instant details (word hover)
- [ ] Get all the bible texts working properly
  - [x] Split GNT text into word, pre, post (separate out punctuation)
  - [ ] Add HOT
  - [ ] Add ESV
  - [x] Add LXX

**Future features**

- [ ] Choose text modal search functionality
- [ ] Search
  - [ ] Functionality
  - [ ] Modal
- [ ] Functions (format the text on screen using functions)
  - [ ] Functionality
  - [ ] Modal
  - [ ] Saved functions
- [ ] Settings window
  - [ ] General (startup)
  - [ ] Appearence (theme, font/size)
  - [ ] Shortcuts
  - [ ] Formatting (e.g. default colours)
  - [ ] Search (e.g. ranges, default text)
  - [ ] Instant details (what is display and order?)
  - [ ] Version/updates
- [ ] Auto-download database file if it does not exist on startup
- [ ] Create menubar with full features
- [ ] Resizing window keeps verse at top on screen, not just block
- [ ] Display app version in bottom right?
- [ ] Right click functionality
  - [ ] On word in text
- [ ] Text highlighting
- [ ] Parallel texts
- [ ] Sidebar holds different app states (?)
- [ ] Add footnoted variant spellings/meanings
- [ ] Add conjoin word highlighting for instant details
- [ ] Add lexicon for more information on instant details
- [ ] Add morph code explanations

**Performance**

- [ ] Split `GetScriptureSections` request into multiple batches for larger ranges that perform concurrently
- [ ] Potentially increase virtualiser speed for resizes / first load

**Bugs**

- [ ] If items changes then virtualiser doesn't invalidate positioning data
- [ ] Virtualiser goto function does not put you in correct scroll for currentRef to be right (e.g. goto Mark 1 shows Matt 28 in currentRef label)
- [ ] Does not currently handle missing verses well (e.g. goto)
- [ ] Should hide chapters with no verses (e.g. LXX Proverbs 25-29)
