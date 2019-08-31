package vugu

// GENERATED FILE, DO NOT EDIT!  See renderer-js-script-maker.go

const jsHelperScript = "\n(function() {\n\n\tif (window.vuguRender) { return; } // only once\n\n    const opcodeEnd = 0         // no more instructions in this buffer\n    // const opcodeClearRefmap = 1 // clear the reference map, all following instructions must not reference prior IDs\n    const opcodeClearEl = 1 // clear the currently selected element\n    // const opcodeSetHTMLRef = 2  // assign ref for html tag\n    // const opcodeSetHeadRef = 3  // assign ref for head tag\n    // const opcodeSetBodyRef = 4  // assign ref for body tag\n    // const opcodeSelectRef = 5   // select element by ref\n\tconst opcodeRemoveOtherAttrs = 5 // remove any elements for the current element that we didn't just set\n    const opcodeSetAttrStr = 6  // assign attribute string to the current selected element\n    const opcodeSelectMountPoint = 7 // selects the mount point element and pushes to the stack - the first time by selector but every subsequent time it will reuse the element from before (because the selector may not match after it's been synced over, it's id etc), also make sure it's of this element name and recreate if so\n\t// const opcodePicardFirstChildElement = 8  // ensure an element first child and push onto element stack\n\t// const opcodePicardFirstChildText    = 9  // ensure a text first child and push onto element stack\n\t// const opcodePicardFirstChildComment = 10 // ensure a comment first child and push onto element stack\n\t// const opcodeSelectParent                   = 11 // pop from the element stack\n\t// const opcodePicardFirstChild = 12  // ensure an element first child and push onto element stack\n\n    const opcodeMoveToFirstChild     = 20 // move node selection to first child (doesn't have to exist)\n\tconst opcodeSetElement           = 21 // assign current selected node as an element of the specified type\n\t// const opcodeSetElementAttr       = 22 // set attribute on current element\n\tconst opcodeSetText              = 23 // assign current selected node as text with specified content\n\tconst opcodeSetComment           = 24 // assign current selected node as comment with specified content\n\tconst opcodeMoveToParent         = 25 // move node selection to parent\n\tconst opcodeMoveToNextSibling    = 26 // move node selection to next sibling (doesn't have to exist)\n\tconst opcodeRemoveOtherEventListeners  = 27 // remove all event listeners from currently selected element that were not just set\n\tconst opcodeSetEventListener     = 28 // assign event listener to currently selected element\n    const opcodeSetInnerHTML         = 29 // set the innerHTML for an element\n\n    const opcodeSetCSSTag            = 30 // write a CSS (style or link) tag\n\tconst opcodeRemoveOtherCSSTags   = 31 // remove any CSS tags that have not been written since the last call\n\tconst opcodeSetJSTag             = 32 // write a JS (script) tag\n    const opcodeRemoveOtherJSTags    = 33 // remove any JS tags that have not been written since the last call\n    \n    const opcodeSetProperty          = 35 // assign a JS property to the current element\n\n\n    // Decoder provides our binary decoding.\n    // Using a class because that's what all the cool JS kids are doing these days.\n    class Decoder {\n\n        constructor(dataView, offset) {\n            this.dataView = dataView;\n            this.offset = offset || 0;\n            return this;\n        }\n\n        // readUint8 reads a single byte, 0-255\n        readUint8() {\n            var ret = this.dataView.getUint8(this.offset);\n            this.offset++;\n            return ret;\n        }\n\n        // readRefToString reads a 64-bit unsigned int ref but returns it as a hex string\n        readRefToString() {\n            // read in two 32-bit parts, BigInt is not yet well supported\n            var ret = this.dataView.getUint32(this.offset).toString(16).padStart(8, \"0\") +\n                this.dataView.getUint32(this.offset + 4).toString(16).padStart(8, \"0\");\n            this.offset += 8;\n            return ret;\n        }\n\n        // readString is 4 bytes length followed by utf chars\n        readString() {\n            var len = this.dataView.getUint32(this.offset);\n            var ret = utf8decoder.decode(new DataView(this.dataView.buffer, this.dataView.byteOffset + this.offset + 4, len));\n            this.offset += len + 4;\n            return ret;\n        }\n\n    }\n\n    let utf8decoder = new TextDecoder();\n\n    window.vuguGetActiveEvent = function() {\n        let state = window.vuguState || {}; window.vuguState = state;\n        return state.activeEvent;\n    }\n    window.vuguGetActiveEventTarget = function() {\n        let state = window.vuguState || {}; window.vuguState = state;\n        return state.activeEvent && state.activeEvent.target;\n    }\n    window.vuguGetActiveEventCurrentTarget = function() {\n        let state = window.vuguState || {}; window.vuguState = state;\n        return state.activeEvent && state.activeEvent.currentTarget;\n    }\n    window.vuguActiveEventPreventDefault = function() {\n        let state = window.vuguState || {}; window.vuguState = state;\n        if (state.activeEvent && state.activeEvent.preventDefault) {\n            state.activeEvent.preventDefault();\n        }\n    }\n    window.vuguActiveEventStopPropagation = function() {\n        let state = window.vuguState || {}; window.vuguState = state;\n        if (state.activeEvent && state.activeEvent.stopPropagation) {\n            state.activeEvent.stopPropagation();\n        }\n    }\n\n\twindow.vuguSetEventHandlerAndBuffer = function(eventHandlerFunc, eventBuffer) { \n\t\tlet state = window.vuguState || {};\n        window.vuguState = state;\n        state.eventBuffer = eventBuffer;\n        state.eventBufferView = new DataView(eventBuffer.buffer, eventBuffer.byteOffset, eventBuffer.byteLength);\n        state.eventHandlerFunc = eventHandlerFunc;\n    }\n\n\twindow.vuguRender = function(buffer) { \n        \n        // NOTE: vuguRender must not automatically reset anything between calls.\n        // Since a series of instructions might get cut off due to buffer end, we\n        // need to be able to just pick right up with the next call where we left off.\n        // The caller decides when to reset things by sending the appropriate\n        // instruction(s).\n\n\t\tlet state = window.vuguState || {};\n\t\twindow.vuguState = state;\n\n\t\tconsole.log(\"vuguRender called\");\n\n        let textEncoder = new TextEncoder();\n\n\t\tlet bufferView = new DataView(buffer.buffer, buffer.byteOffset, buffer.byteLength);\n\n        var decoder = new Decoder(bufferView, 0);\n        \n        // state.refMap = state.refMap || {};\n        // state.curRef = state.curRef || \"\"; // current reference number (as a hex string)\n        // state.curRefEl = state.curRefEl || null; // current reference element\n        // state.elStack = state.elStack || []; // stack of elements as we traverse the DOM tree\n\n        // mount point element\n        state.mountPointEl = state.mountPointEl || null; \n\n        // currently selected element\n        state.el = state.el || null;\n\n        // specifies a \"next\" move for the current element, if used it must be followed by\n        // one of opcodeSetElement, opcodeSetText, opcodeSetComment, which will create/replace/use existing\n        // the element and put it in \"el\".  The point is this allow us to select nodes that may\n        // not exist yet, knowing that the next call will specify what that node is.  It's more complex here\n        // but makes it easier to generate instructions while walking a DOM tree.\n        // Value is one of \"first_child\", \"next_sibling\"\n        // (Parents always exist and so doesn't use this mechanism.)\n        state.nextElMove = state.nextElMove || null;\n\n        // keeps track of attributes that are being set on the current element, so we can remove any extras\n        state.elAttrNames = state.elAttrNames || {};\n\n        // map of positionID -> array of listener spec and handler function, for all elements\n        state.eventHandlerMap = state.eventHandlerMap || {};\n    \n        // keeps track of event listeners that are being set on the current element, so we can remvoe any extras\n        state.elEventKeys = state.elEventKeys || {};\n\n        instructionLoop: while (true) {\n\n            let opcode = decoder.readUint8();\n            \n            try {\n\n                console.log(\"processing opcode\", opcode);\n                // console.log(\"test_span_id: \", document.querySelector(\"#test_span_id\"));\n\n                switch (opcode) {\n\n                    case opcodeEnd: {\n                        break instructionLoop;\n                    }\n        \n                    case opcodeClearEl: {\n                        state.el = null;\n                        state.nextElMove = null;\n                        break;\n                    }\n\n                    case opcodeSetProperty: {\n                        let el = state.el;\n                        if (!el) {\n                            return \"opcodeSetProperty: no current reference\";\n                        }\n                        let propName = decoder.readString();\n                        let propValueJSON = decoder.readString();\n                        el[propName] = JSON.parse(propValueJSON);\n                        break;\n                    }                    \n            \n                    case opcodeSetAttrStr: {\n                        let el = state.el;\n                        if (!el) {\n                            return \"opcodeSetAttrStr: no current reference\";\n                        }\n                        let attrName = decoder.readString();\n                        let attrValue = decoder.readString();\n                        el.setAttribute(attrName, attrValue);\n                        state.elAttrNames[attrName] = true;\n                        // console.log(\"setting attr\", attrName, attrValue, el)\n                        break;\n                    }\n\n                    case opcodeSelectMountPoint: {\n                        \n                        state.elAttrNames = {}; // reset attribute list\n                        state.elEventKeys = {};\n\n                        // select mount point using selector or if it was done earlier re-use the one from before\n                        let selector = decoder.readString();\n                        let nodeName = decoder.readString();\n                        // console.log(\"GOT HERE selector,nodeName = \", selector, nodeName);\n                        // console.log(\"state.mountPointEl\", state.mountPointEl);\n                        if (state.mountPointEl) {\n                            console.log(\"opcodeSelectMountPoint: state.mountPointEl already exists, using it\", state.mountPointEl, \"parent is\", state.mountPointEl.parentNode);\n                            state.el = state.mountPointEl;\n                            // state.elStack.push(state.mountPointEl);\n                        } else {\n                            console.log(\"opcodeSelectMountPoint: state.mountPointEl does not exist, using selector to find it\", selector);\n                            let el = document.querySelector(selector);\n                            if (!el) {\n                                throw \"mount point selector not found: \" + selector;\n                            }\n                            state.mountPointEl = el;\n                            // state.elStack.push(el);\n                            state.el = el;\n                        }\n\n                        let el = state.el;\n\n                        // make sure it's the right element name and replace if not\n                        if (el.nodeName.toUpperCase() != nodeName.toUpperCase()) {\n\n                            let newEl = document.createElement(nodeName);\n                            el.parentNode.replaceChild(newEl, el);\n\n                            state.mountPointEl = newEl;\n                            el = newEl;\n\n                        }\n\n                        state.el = el;\n\n                        state.nextElMove = null;\n\n                        break;\n                    }\n\n                    // remove any elements for the current element that we didn't just set\n                    case opcodeRemoveOtherAttrs: {\n\n                        if (!state.el) {\n                            throw \"no element selected\";\n                        }\n\n                        if (state.nextElMove) {\n                            throw \"cannot call opcodeRemoveOtherAttrs when nextElMove is set\";\n                        }\n\n                        // build a list of attribute names to remove\n                        let rmAttrNames = [];\n                        for (let i = 0; i < state.el.attributes.length; i++) {\n                            if (!state.elAttrNames[state.el.attributes[i].name]) {\n                                rmAttrNames.push(state.el.attributes[i].name);\n                            }\n                        }\n\n                        // remove them\n                        for (let i = 0; i < rmAttrNames.length; i++) {\n                            state.el.attributes.removeNamedItem(rmAttrNames[i]);\n                        }\n\n                        break;\n                    }\n\n                    // move node selection to parent\n                    case opcodeMoveToParent: {\n\n                        // if first_child is next move then we just unset this\n                        if (state.nextElMove == \"first_child\") {\n                            state.nextElMove = null;\n                        } else {\n                            // otherwise we actually move and also reset nextElMove\n                            state.el = state.el.parentNode;\n                            state.nextElMove = null;\n                        }\n\n                        break;\n                    }\n\n                    // move node selection to first child (doesn't have to exist)\n                    case opcodeMoveToFirstChild: {\n\n                        // if a next move already set, then we need to execute it before we can do this\n                        if (state.nextElMove) {\n                            if (state.nextElMove == \"first_child\") {\n                                state.el = state.el.firstChild;\n                                if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                            } else if (state.nextElMove == \"next_sibling\") {\n                                state.el = state.el.nextSibling;\n                                if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                            }\n                            state.nextElMove = null;\n                        }\n\n                        if (!state.el) { throw \"must have current selection to use opcodeMoveToFirstChild\"; }\n                        state.nextElMove = \"first_child\";\n\n                        break;\n                    }\n                    \n                    // move node selection to next sibling (doesn't have to exist)\n                    case opcodeMoveToNextSibling: {\n\n                        // if a next move already set, then we need to execute it before we can do this\n                        if (state.nextElMove) {\n                            if (state.nextElMove == \"first_child\") {\n                                state.el = state.el.firstChild;\n                                if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                            } else if (state.nextElMove == \"next_sibling\") {\n                                state.el = state.el.nextSibling;\n                                if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                            }\n                            state.nextElMove = null;\n                        }\n\n                        if (!state.el) { throw \"must have current selection to use opcodeMoveToNextSibling\"; }\n                        state.nextElMove = \"next_sibling\";\n\n                        break;\n                    }\n                    \n                    // assign current selected node as an element of the specified type\n                    case opcodeSetElement: {\n                        \n                        let nodeName = decoder.readString();\n\n                        this.console.log(\"opcodeSetElement for \",\n                            \"nodeName=\", nodeName, \n                            \", state.el=\", state.el, \n                            \", state.nextElMove=\", state.nextElMove);\n\n                        state.elAttrNames = {};\n                        state.elEventKeys = {};\n\n                        // handle nextElMove cases\n\n                        if (state.nextElMove == \"first_child\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.firstChild;\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                newEl = document.createElement(nodeName);\n                                state.el.appendChild(newEl);\n                                state.el = newEl;\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.nextSibling;\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                newEl = document.createElement(nodeName);\n                                // console.log(\"HERE1\", state.el);\n                                // state.el.insertAdjacentElement(newEl, 'afterend');\n                                state.el.parentNode.appendChild(newEl);\n                                state.el = newEl;\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove) {\n                            throw \"bad state.nextElMove value: \" + state.nextElMove;\n                        }\n\n                        // if we get here we need to verify that state.el is in fact an element of the right type\n                        // and replace if not\n\n                        if (state.el.nodeType != 1 || state.el.nodeName.toUpperCase() != nodeName.toUpperCase()) {\n\n                            let newEl = document.createElement(nodeName);\n                            // throw \"stopping here\";\n                            state.el.parentNode.replaceChild(newEl, state.el);\n                            state.el = newEl;\n\n                        }\n\n                        break;\n                    }\n\n                    // assign current selected node as text with specified content\n                    case opcodeSetText: {\n\n                        let content = decoder.readString();\n\n                        // console.log(\"in opcodeSetText 1\");\n\n                        // handle nextElMove cases\n\n                        if (state.nextElMove == \"first_child\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.firstChild;\n                            // console.log(\"in opcodeSetText 2\");\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                let newEl = document.createTextNode(content);\n                                state.el.appendChild(newEl);\n                                state.el = newEl;\n                                // console.log(\"in opcodeSetText 3\");\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.nextSibling;\n                            // console.log(\"in opcodeSetText 4\");\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                let newEl = document.createTextNode(content);\n                                // state.el.insertAdjacentElement(newEl, 'afterend');\n                                state.el.parentNode.appendChild(newEl);\n                                state.el = newEl;\n                                // console.log(\"in opcodeSetText 5\");\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove) {\n                            throw \"bad state.nextElMove value: \" + state.nextElMove;\n                        }\n\n                        // if we get here we need to verify that state.el is in fact a node of the right type\n                        // and with right content and replace if not\n                        // console.log(\"in opcodeSetText 6\");\n\n                        if (state.el.nodeType != 3) {\n\n                            let newEl = document.createTextNode(content);\n                            state.el.parentNode.replaceChild(newEl, state.el);\n                            state.el = newEl;\n                            // console.log(\"in opcodeSetText 7\");\n\n                        } else {\n                            // console.log(\"in opcodeSetText 8\");\n                            state.el.textContent = content;\n                        }\n                        // console.log(\"in opcodeSetText 9\");\n\n                        break;\n                    }\n\n                    // assign current selected node as comment with specified content\n                    case opcodeSetComment: {\n                        \n                        let content = decoder.readString();\n\n                        // handle nextElMove cases\n\n                        if (state.nextElMove == \"first_child\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.firstChild;\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                let newEl = document.createComment(content);\n                                state.el.appendChild(newEl);\n                                state.el = newEl;\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.nextElMove = null;\n                            let newEl = state.el.nextSibling;\n                            if (newEl) { \n                                state.el = newEl; \n                                // fall through to verify state.el is correct below\n                            } else {\n                                let newEl = document.createComment(content);\n                                // state.el.insertAdjacentElement(newEl, 'afterend');\n                                state.el.parentNode.appendChild(newEl);\n                                state.el = newEl;\n                                break; // we're done here, since we just created the right element\n                            }\n                        } else if (state.nextElMove) {\n                            throw \"bad state.nextElMove value: \" + state.nextElMove;\n                        }\n\n                        // if we get here we need to verify that state.el is in fact a node of the right type\n                        // and with right content and replace if not\n\n                        if (state.el.nodeType != 8) {\n\n                            let newEl = document.createComment(content);\n                            state.el.parentNode.replaceChild(newEl, state.el);\n                            state.el = newEl;\n\n                        } else {\n                            state.el.textContent = content;\n                        }\n\n                        break;\n                    }\n\n                    case opcodeSetInnerHTML: {\n\n                        let html = decoder.readString();\n\n                        if (!state.el) { throw \"opcodeSetInnerHTML must have currently selected element\"; }\n                        if (state.nextElMove) { throw \"opcodeSetInnerHTML nextElMove must not be set\"; }\n                        if (state.el.nodeType != 1) { throw \"opcodeSetInnerHTML currently selected element expected nodeType 1 but has: \" + state.el.nodeType; }\n\n                        state.el.innerHTML = html;\n\n                        break;\n                    }\n\n                    // remove all event listeners from currently selected element that were not just set\n                    case opcodeRemoveOtherEventListeners: {\n                        this.console.log(\"opcodeRemoveOtherEventListeners\");\n\n                        let positionID = decoder.readString();\n\n                        // look at all registered events for this positionID\n                        let emap = state.eventHandlerMap[positionID] || {};\n                        // for any that we didn't just set, remove them\n                        let toBeRemoved = [];\n                        for (let k in emap) {\n                            if (!state.elEventKeys[k]) {\n                                toBeRemoved.push(k);\n                            }\n                        }\n\n                        // for each one that was missing, we remove from emap and call removeEventListener\n                        for (let i = 0; i < toBeRemoved.length; i++) {\n                            let f = emap[k];\n                            let k = toBeRemoved[i];\n                            let kparts = k.split(\"|\");\n                            state.el.removeEventListener(kparts[0], f, {capture:!!kparts[1], passive:!!kparts[2]});\n                            delete emap[k];\n                        }\n\n                        // if emap is empty now, remove the entry from eventHandlerMap altogether\n                        if (Object.keys(emap).length == 0) {\n                            delete state.eventHandlerMap[positionID];\n                        } else {\n                            state.eventHandlerMap[positionID] = emap;\n                        }\n\n                        break;\n                    }\n                \n                    // assign event listener to currently selected element\n                    case opcodeSetEventListener: {\n                        let positionID = decoder.readString();\n                        let eventType = decoder.readString();\n                        let capture = decoder.readUint8();\n                        let passive = decoder.readUint8();\n\n                        if (!state.el) {\n                            throw \"must have state.el set in order to call opcodeSetEventListener\";\n                        }\n\n                        var eventKey = eventType + \"|\" + (capture?\"1\":\"0\") + \"|\" + (passive?\"1\":\"0\");\n                        state.elEventKeys[eventKey] = true;\n\n                        // map of positionID -> map of listener spec and handler function, for all elements\n                        //state.eventHandlerMap\n                        let emap = state.eventHandlerMap[positionID] || {};\n\n                        // register function if not done already\n                        let f = emap[eventKey];\n                        if (!f) {\n                            f = function(event) {\n\n                                // set the active event, so the Go code and call back in and examine it if needed\n                                state.activeEvent = event; \n\n                                let eventObj = {};\n                                // console.log(event);\n                                for (let i in event) {\n                                    let itype = typeof(event[i]);\n                                    // copy primitive values directly\n                                    if ((itype == \"boolean\" || itype == \"number\" || itype == \"string\") && true/*event.hasOwnProperty(i)*/) {\n                                        eventObj[i] = event[i];\n                                    }\n                                }\n\n                                // also do the same for anything in \"target\"\n                                if (event.target) {\n                                    eventObj.target = {};\n                                    let et = event.target;\n                                    for (let i in et) {\n                                        let itype = typeof(et[i]);\n                                        if ((itype == \"boolean\" || itype == \"number\" || itype == \"string\") && true/*et.hasOwnProperty(i)*/) {\n                                            eventObj.target[i] = et[i];\n                                        }\n                                    }\n                                }\n                                \n                                // console.log(eventObj);\n                                // console.log(JSON.stringify(eventObj));\n\n                                let fullJSON = JSON.stringify({\n                                    \n                                    // include properties from event registration\n                                    position_id: positionID,\n                                    event_type: eventType,\n                                    capture: !!capture,\n                                    passive: !!passive,\n\n                                    // the event object data as extracted above\n                                    event_summary: eventObj,\n\n                                });\n\n                                // console.log(state.eventBuffer);\n\n                                // write JSON to state.eventBuffer with zero char as termination\n\n                                \n                                let encodeResultBuffer = textEncoder.encode(fullJSON);\n                                //console.log(\"encodeResult\", encodeResult);\n                                state.eventBuffer.set(encodeResultBuffer, 4); // copy encoded string to event buffer\n                                // now write length using DataView as uint32\n                                state.eventBufferView.setUint32(0, encodeResultBuffer.byteLength - encodeResultBuffer.byteOffset);\n\n                                // let result = textEncoder.encodeInto(fullJSON, state.eventBuffer);\n                                // let eventBufferDataView = new DataView(state.eventBuffer.buffer, state.eventBuffer.byteOffset, state.eventBuffer.byteLength);\n                                // eventBufferDataView.setUint8(result.written, 0);\n\n                                // write length after, since only now do we know the final length\n                                // state.eventBufferView.setUint32(0, result.written);\n\n                                // serialize event into the event buffer, somehow,\n                                // and keep track of the target element, also consider grabbing\n                                // the value or relevant properties as appropriate for form things\n                                \n                                state.eventHandlerFunc.call(null); // call with null this avoid unnecessary js.Value reference\n\n                                // unset the active event\n                                state.activeEvent = null;\n                            };    \n                            emap[eventKey] = f;\n\n                            // this.console.log(\"addEventListener\", eventType);\n                            state.el.addEventListener(eventType, f, {capture:capture, passive:passive});\n                        }\n\n                        state.eventHandlerMap[positionID] = emap;\n\n                        this.console.log(\"opcodeSetEventListener\", positionID, eventType, capture, passive);\n                        break;\n                    }\n\n                    case opcodeSetCSSTag: {\n\n                        let elementName = decoder.readString();\n                        let textContent = decoder.readString();\n                        let attrPairsLen = decoder.readUint8();\n                        if (attrPairsLen % 2 != 0) {\n                            throw \"attrPairsLen is odd number: \" + attrPairsLen;\n                        }\n                        // loop over one key/value pair at a time and put them in a map\n                        var attrMap = {};\n                        for (let i = 0; i < attrPairsLen; i += 2) {\n                            let key = decoder.readString();\n                            let val = decoder.readString();\n                            attrMap[key] = val;\n                        }\n\n                        this.console.log(\"got opcodeSetCSSTag: elementName=\", elementName, \"textContent=\", textContent, \"attrMap=\", attrMap)\n                        \n                        state.elCSSTagsSet = state.elCSSTagsSet || []; // ensure state.elCSSTagsSet is set to empty array if not already set\n\n                        // let elementNameUC = elementName.toUpperCase();\n                        let thisTagKey = textContent;\n                        if (elementName == \"link\") {\n                            thisTagKey = attrMap[\"href\"];\n                        }\n\n                        if (thisTagKey == \"\") { // nothing to do in this case\n                            this.console.log(\"element\", elementName, \"ignored due to empty key\");\n                            break;\n                        }\n\n                        // TODO: \n                        // * find all tags that have the same element type (link or style)\n                        // * for each one for style use textContent as key, for link use url\n                        // * see if matching tag already exists\n                        // * if it has vuguCreated==true on it, then add to map of css tags set, else ignore\n                        // * if no matching tag then create and set vuguCreated=true, add to map of css tags set\n\n                        let foundTag = null;\n                        this.document.querySelectorAll(elementName).forEach(cssEl => {\n                            let cssElKey;\n                            if (elementName == \"style\") {\n                                cssElKey = cssEl.textContent;\n                            } else /* elementName == \"link\" */ {\n                                cssElKey = cssEl.href;\n                            }\n                            \n                            if (thisTagKey == cssElKey) { // textContent or href as appropriate is used to determine \"sameness\"\n                                foundTag = cssEl;\n                            }\n                        });\n\n                        // could not find it, create\n                        if (!foundTag) {\n                            let cTag = this.document.createElement(elementName);\n                            for (let k in attrMap) {\n                                cTag.setAttribute(k, attrMap[k]);\n                            }\n                            cTag.vuguCreated = true; // so we know that we created this, as opposed to it already having been on the page\n                            this.document.head.appendChild(cTag); // add to end of head\n                            state.elCSSTagsSet.push(cTag); // add to elCSSTagsSet for use in opcodeRemoveOtherCSSTags\n                        } else {\n                            // if we did find it, we need to push to state.elCSSTagsSet to tell opcodeRemoveOtherCSSTags not to remove it\n                            state.elCSSTagsSet.push(foundTag);\n                        }\n\n                        break;\n                    }\n                    case opcodeRemoveOtherCSSTags: {\n\n                        this.console.log(\"got opcodeRemoveOtherCSSTags\");\n\n                        // any link or style tag in doc that has vuguCreated==true and is not in css tags set map gets removed\n\n                        state.elCSSTagsSet = state.elCSSTagsSet || [];\n\n                        this.document.querySelectorAll('style,link').forEach(cssEl => {\n\n                            // ignore any not created by vugu\n                            if (!cssEl.vuguCreated) { return; }\n\n                            // ignore if in elCSSTagsSet\n                            if (state.elCSSTagsSet.findIndex(el => el==cssEl) >= 0) { return; }\n\n                            // if we got here, we remove the tag\n                            cssEl.parentNode.removeChild(cssEl);\n                        });\n\n                        state.elCSSTagsSet = null; // clear this out so it gets reinitialized the next time opcodeSetCSSTag or this opcode is used\n\n                        break;\n                    }\n\n                    default: {\n                        console.error(\"found invalid opcode\", opcode);\n                        return;\n                    }\n                }\n\n            } catch (e) {\n                this.console.log(\"Error during instruction loop. Data opcode=\", opcode, \n                    \", state.el=\", state.el, \n                    \", state.nextElMove=\", state.nextElMove, \n                    \", with error: \", e)\n                throw e;\n            }\n\n\n\t\t}\n\n\t}\n\n})()\n"