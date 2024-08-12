import React from 'react';
import { SearchInput, SelectOption } from '@patternfly/react-core';

import { SearchFilter } from 'types/search';
import { SelectedEntity } from './EntitySelector';
import { SelectedAttribute } from './AttributeSelector';
import { CompoundSearchFilterConfig, OnSearchPayload } from '../types';
import {
    conditionMap,
    dateConditionMap,
    ensureConditionDate,
    ensureConditionNumber,
    ensureString,
    ensureStringArray,
    getAttribute,
    getEntity,
    isSelectType,
} from '../utils/utils';

import CheckboxSelect from './CheckboxSelect';
import ConditionNumber from './ConditionNumber';
import SearchFilterAutocomplete from './SearchFilterAutocomplete';
import ConditionDate from './ConditionDate';

export type InputFieldValue =
    | string
    | number
    | undefined
    | string[]
    | { condition: string; number: number }
    | { condition: string; date: string };
export type InputFieldOnChange = (value: InputFieldValue) => void;

export type CompoundSearchFilterInputFieldProps = {
    selectedEntity: SelectedEntity;
    selectedAttribute: SelectedAttribute;
    value: InputFieldValue;
    searchFilter: SearchFilter;
    additionalContextFilter?: SearchFilter;
    onSearch: ({ action, category, value }: OnSearchPayload) => void;
    onChange: InputFieldOnChange;
    config: CompoundSearchFilterConfig;
};

function CompoundSearchFilterInputField({
    selectedEntity,
    selectedAttribute,
    value,
    searchFilter,
    additionalContextFilter,
    onSearch,
    onChange,
    config,
}: CompoundSearchFilterInputFieldProps) {
    if (!selectedEntity || !selectedAttribute) {
        return null;
    }

    const entity = getEntity(config, selectedEntity);
    const attribute = getAttribute(config, selectedEntity, selectedAttribute);

    if (!attribute) {
        return null;
    }

    if (attribute.inputType === 'text') {
        const textLabel = `Filter results by ${attribute.filterChipLabel}`;
        return (
            <SearchInput
                aria-label={textLabel}
                placeholder={textLabel}
                value={ensureString(value)}
                onChange={(_event, _value) => onChange(_value)}
                onSearch={(_event, _value) => {
                    onSearch({
                        action: 'ADD',
                        category: attribute.searchTerm,
                        value: _value,
                    });
                    onChange('');
                }}
                onClear={() => onChange('')}
                submitSearchButtonLabel="Apply text input to search"
            />
        );
    }
    if (attribute.inputType === 'date-picker') {
        return (
            <ConditionDate
                value={ensureConditionDate(value)}
                onChange={(newValue) => {
                    onChange(newValue);
                }}
                onSearch={(newValue) => {
                    const { condition, date } = newValue;
                    onSearch({
                        action: 'ADD',
                        category: attribute.searchTerm,
                        value: `${dateConditionMap[condition]}${date}`,
                    });
                    onChange({ ...newValue, date: '' });
                }}
            />
        );
    }
    if (attribute.inputType === 'condition-number') {
        return (
            <ConditionNumber
                value={ensureConditionNumber(value)}
                onChange={(newValue) => {
                    onChange(newValue);
                }}
                onSearch={(newValue) => {
                    const { condition, number } = newValue;
                    onChange(newValue);
                    onSearch({
                        action: 'ADD',
                        category: attribute.searchTerm,
                        value: `${conditionMap[condition]}${number}`,
                    });
                }}
            />
        );
    }
    if (entity && entity.searchCategory && attribute.inputType === 'autocomplete') {
        const { searchCategory } = entity;
        const { searchTerm, filterChipLabel } = attribute;
        const textLabel = `Filter results by ${filterChipLabel}`;
        return (
            <SearchFilterAutocomplete
                searchCategory={searchCategory}
                searchTerm={searchTerm}
                value={ensureString(value)}
                onChange={(newValue) => {
                    onChange(newValue);
                }}
                onSearch={(newValue) => {
                    onSearch({
                        action: 'ADD',
                        category: attribute.searchTerm,
                        value: newValue,
                    });
                    onChange('');
                }}
                textLabel={textLabel}
                searchFilter={searchFilter}
                additionalContextFilter={additionalContextFilter}
            />
        );
    }
    if (isSelectType(attribute)) {
        const attributeLabel = attribute.displayName;
        const selectOptions = attribute.inputProps.options;
        const { searchTerm } = attribute;
        const selection = ensureStringArray(searchFilter?.[searchTerm]);

        return (
            <CheckboxSelect
                selection={selection}
                onChange={(checked, _value) => {
                    onChange(value);
                    onSearch({
                        action: checked ? 'ADD' : 'REMOVE',
                        category: attribute.searchTerm,
                        value: _value,
                    });
                }}
                ariaLabelMenu={`Filter by ${attributeLabel} select menu`}
                toggleLabel={`Filter by ${attributeLabel}`}
            >
                {selectOptions.length !== 0 ? (
                    selectOptions.map((option) => {
                        return (
                            <SelectOption
                                key={option.value}
                                hasCheckbox
                                value={option.value}
                                isSelected={selection.includes(option.value)}
                            >
                                {option.label}
                            </SelectOption>
                        );
                    })
                ) : (
                    <SelectOption isDisabled>No options available</SelectOption>
                )}
            </CheckboxSelect>
        );
    }
    return <div />;
}

export default CompoundSearchFilterInputField;
