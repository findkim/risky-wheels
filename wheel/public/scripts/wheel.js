// Rending and annimation of spin wheel

// Helpers
function shuffle(o) {
	for ( var j, x, i = o.length; i; j = parseInt(Math.random() * i), x = o[--i], o[i] = o[j], o[j] = x)
		;
	return o;
};

String.prototype.hashCode = function(){
	// See http://www.cse.yorku.ca/~oz/hash.html		
	var hash = 5381;
	for (i = 0; i < this.length; i++) {
		char = this.charCodeAt(i);
		hash = ((hash<<5)+hash) + char;
		hash = hash & hash; // Convert to 32bit integer
	}
	return hash;
}

Number.prototype.mod = function(n) {
	return ((this%n)+n)%n;
}

// Wheel
var wheel = {
	timerHandle : 0,
	timerDelay : 33,
	angleCurrent : 0,
	angleDelta : 0,
	size : 290,
	canvasContext : null,
	colors : [ '#fd7c66', '#eabc20', '#3b9fe8', '#53ff45', '#e2a3ff'],
	segments : [],
	seg_colors : [], // Cache of segments to colors
	maxSpeed : Math.PI / 8,
	upTime : 1000, // How long to spin up for (in ms)
	downTime : 5000, // How long to slow down for (in ms)
	spinStart : 0,
	frames : 0,
	centerX : 300,
	centerY : 300,

	spin : function() {
		// Start the wheel only if it's not already spinning
		if (wheel.timerHandle == 0) {
			wheel.spinStart = new Date().getTime();
			wheel.maxSpeed = Math.PI / (16 + Math.random()); // Randomly vary how hard the spin is
			wheel.frames = 0;
			// wheel.sound.play();
			wheel.timerHandle = setInterval(wheel.onTimerTick, wheel.timerDelay);
		}
	},

	onTimerTick : function() {
		wheel.frames++;
		wheel.draw();

		var duration = (new Date().getTime() - wheel.spinStart);
		var progress = 0;
		var finished = false;

		if (duration < wheel.upTime) {
			progress = duration / wheel.upTime;
			wheel.angleDelta = wheel.maxSpeed
					* Math.sin(progress * Math.PI / 2);
		} else {
			progress = duration / wheel.downTime;
			wheel.angleDelta = wheel.maxSpeed
					* Math.sin(progress * Math.PI / 2 + Math.PI / 2);
			if (progress >= 1)
				finished = true;
		}

		wheel.angleCurrent += wheel.angleDelta;
		while (wheel.angleCurrent >= Math.PI * 2)
			// Keep the angle in a reasonable range
			wheel.angleCurrent -= Math.PI * 2;

		if (finished) {
			clearInterval(wheel.timerHandle);
			wheel.timerHandle = 0;
			wheel.angleDelta = 0;
		}
	},

	init : function(optionList) {
		try {
			document.fonts.load('2em "Burbank Regular"').then(function(){
				wheel.initWheel();
				// wheel.initAudio();
				wheel.initCanvas();
				wheel.draw();
			});

			$.extend(wheel, optionList);

		} catch (exceptionData) {
			alert('Wheel is not loaded ' + exceptionData);
		}
	},

	// initAudio : function() {
	// 	var sound = document.createElement('audio');
	// 	sound.setAttribute('src', 'wheel.mp3');
	// 	wheel.sound = sound;
	// },

	initCanvas : function() {
		// var canvas = $('#wheel #canvas').get(0);
		var canvas = document.getElementById("wheel-canvas")
		canvas.addEventListener("click", wheel.spin, false);
		wheel.canvasContext = canvas.getContext("2d");
	},

	initWheel : function() {
		shuffle(wheel.colors);
	},

	// Called when segments have changed
	update : function() {
		// Ensure we start mid way on a item
		//var r = Math.floor(Math.random() * wheel.segments.length);
		var r = 0;
		wheel.angleCurrent = ((r + 0.5) / wheel.segments.length) * Math.PI * 2;

		var segments = wheel.segments;
		var len      = segments.length;
		var colors   = wheel.colors;
		var colorLen = colors.length;

		// Generate a color cache (so we have consistant coloring)
		var seg_color = new Array();
		for (var i = 0; i < len; i++)
			seg_color.push( colors [ segments[i].hashCode().mod(colorLen) ] );
		wheel.seg_color = seg_color;
		wheel.draw();
	},

	draw : function() {
		wheel.clear();
		wheel.drawWheel();
		wheel.drawNeedle();
	},

	clear : function() {
		var ctx = wheel.canvasContext;
		ctx.clearRect(0, 0, 1000, 800);
	},

	drawNeedle : function() {
		var ctx = wheel.canvasContext;
		var centerX = wheel.centerX;
		var centerY = wheel.centerY;
		var size = wheel.size;

		ctx.lineWidth = 1;
		ctx.strokeStyle = '#ffffff';
		ctx.fileStyle = '#ffffff';

		ctx.beginPath();

		ctx.moveTo(centerX + size - 40, centerY);
		ctx.lineTo(centerX + size + 20, centerY - 10);
		ctx.lineTo(centerX + size + 20, centerY + 10);
		ctx.closePath();

		ctx.stroke();
		ctx.fill();

		// Which segment is being pointed to?
		var i = wheel.segments.length - Math.floor((wheel.angleCurrent / (Math.PI * 2))	* wheel.segments.length) - 1;

		// Now draw the winning name
		ctx.textAlign = "left";
		ctx.textBaseline = "middle";
		ctx.fillStyle = '#ffffff';
		ctx.font = "2em Burbank Regular";
		ctx.fillText(wheel.segments[i], centerX + size + 35, centerY);
	},

	drawSegment : function(key, lastAngle, angle) {
		var ctx = wheel.canvasContext;
		var centerX = wheel.centerX;
		var centerY = wheel.centerY;
		var size = wheel.size;

		var segments = wheel.segments;
		var len = wheel.segments.length;
		var colors = wheel.seg_color;

		var value = segments[key];
		
		ctx.save();
		ctx.beginPath();

		// Start in the centre
		ctx.moveTo(centerX, centerY);
		ctx.arc(centerX, centerY, size, lastAngle, angle, false); // Draw a arc around the edge
		ctx.lineTo(centerX, centerY); // Now draw a line back to the centre

		// Clip anything that follows to this area
		//ctx.clip(); // It would be best to clip, but we can double performance without it
		ctx.closePath();

		ctx.fillStyle = colors[key];
		ctx.fill();
		ctx.stroke();

		// Now draw the text
		ctx.save(); // The save ensures this works on Android devices
		ctx.translate(centerX, centerY);
		ctx.rotate((lastAngle + angle) / 2);

		ctx.fillStyle = "#ffffff";
		ctx.fillText(value.substr(0, 20), size / 2 + 20, 0);
		ctx.restore();
	},

	drawWheel : function() {
		var ctx = wheel.canvasContext;

		var angleCurrent = wheel.angleCurrent;
		var lastAngle    = angleCurrent;

		var segments  = wheel.segments;
		var len       = wheel.segments.length;
		var colors    = wheel.colors;
		var colorsLen = wheel.colors.length;

		var centerX = wheel.centerX;
		var centerY = wheel.centerY;
		var size    = wheel.size;

		var PI2 = Math.PI * 2;

		ctx.lineWidth    = 2;
		ctx.strokeStyle  = '#ffffff';
		ctx.textBaseline = "middle";
		ctx.textAlign    = "center";
		ctx.font         = "1.4em Burbank Regular";

		for (var i = 1; i <= len; i++) {
			var angle = PI2 * (i / len) + angleCurrent;
			wheel.drawSegment(i - 1, lastAngle, angle);
			lastAngle = angle;
		}
		// Draw a center circle
		ctx.beginPath();
		ctx.arc(centerX, centerY, 50, 0, PI2, false);
		ctx.closePath();

		ctx.fillStyle   = '#ffffff';
		// ctx.strokeStyle = '#080F31';
		ctx.fill();
		ctx.stroke();

		// Draw outer circle
		ctx.beginPath();
		ctx.arc(centerX, centerY, size, 0, PI2, false);
		ctx.closePath();

		ctx.lineWidth   = 10;
		ctx.strokeStyle = '#080F31';
		ctx.stroke();
	},
}

window.onload = function() {
	wheel.init();

	var segments = new Array();
	$.each($('#locations input:checked'), function(key, cbox) {
		segments.push( cbox.value );
	});

	wheel.segments = segments;
	wheel.update();

	// Hide the address bar (for mobile devices)!
	setTimeout(function() {
		window.scrollTo(0, 1);
	}, 0);
}
