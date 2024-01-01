/**
 * Wasp API
 * REST API for the Wasp node
 *
 * OpenAPI spec version: 0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { StateTransaction } from '../models/StateTransaction';
import { HttpFile } from '../http/http';

export class PublisherStateTransactionItem {
    'lastMessage': StateTransaction;
    'messages': number;
    'timestamp': Date;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "lastMessage",
            "baseName": "lastMessage",
            "type": "StateTransaction",
            "format": ""
        },
        {
            "name": "messages",
            "baseName": "messages",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "timestamp",
            "baseName": "timestamp",
            "type": "Date",
            "format": "date-time"
        }    ];

    static getAttributeTypeMap() {
        return PublisherStateTransactionItem.attributeTypeMap;
    }

    public constructor() {
    }
}

