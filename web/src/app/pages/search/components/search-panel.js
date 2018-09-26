import { h, Component } from 'preact';

import SearchSvg from './../../../../images/search.svg';

import './search-panel.less';

export default class SearchPanel extends Component {
  constructor(props) {
    super(props);

    this.state = {
      searchField: '',
      searchTerm: '',
    };

    this.searchFields = [
      { name: '--All Fields--', value: '' },
      { name: 'Artists', value: 'artist' },
      { name: 'Title', value: 'title' },
      { name: 'Album', value: 'album' },
    ];
  }

  startSearch() {
    const { searchTerm, searchField } = this.state;
    if (searchTerm) {
      this.props.onStartSearch(searchTerm, searchField);
    }
  }

  findDuplicates() {
    this.props.onFindDuplicates();
  }

  handleSearchInputKeyPress(key, value) {
    if (key === 'Enter') {
      this.setState({ searchTerm: value }, this.startSearch);
    }
  }

  renderSearchInput(searchTerm) {
    return (
      <div>
        <input
          class="input"
          type="text"
          placeholder="Enter the search term"
          value={searchTerm}
          onKeyPress={event => this.handleSearchInputKeyPress(event.key, event.target.value)}
          onChange={event => this.setState({ searchTerm: event.target.value })}
        />
      </div>
    );
  }

  renderSearchFields(searchField) {
    return (
      <div>
        <div class="field">
          <div class="control">
            <div class="select">
              <select
                value={searchField}
                onChange={event => this.setState({ searchField: event.target.value })}
              >
                {
                  this.searchFields.map(searchField =>
                    <option value={searchField.value}>{searchField.name}</option>
                  )}
              </select>
            </div>
          </div>
        </div>
      </div>
    );
  }

  renderSearchButton() {
    return (
      <div>
        <a
          class="button is-primary is-rounded"
          onClick={() => this.startSearch()}
        >
          <SearchSvg class="SearchPanel-searchIcon" />
        </a>
      </div>
    );
  }

  renderDuplicatesButton() {
    return (
      <div>
        <a
          class="button is-primary is-rounded"
          onClick={() => this.findDuplicates()}
          title="Duplicates"
        >
          Duplicates
        </a>
      </div>
    );
  }

  render(props, state) {
    return (
      <div class="SearchPanel-container">
        { this.renderSearchInput(state.searchTerm) }
        { this.renderSearchFields(state.searchField) }
        { this.renderSearchButton() }
        { this.renderDuplicatesButton() }
      </div>
    );
  }
}
